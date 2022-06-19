package err

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func TestNewError1(t *testing.T) {
	code := codes.Internal
	reason := "reason"
	message := "message"
	err := New(code, reason, message)

	if code != err.code {
		t.Errorf("\n  got: %s\n  want: %s", err.code, code)
		return
	}
	if reason != err.reason {
		t.Errorf("\n  got: %s\n  want: %s", err.reason, reason)
		return
	}
	if message != err.message {
		t.Errorf("\n  got: %s\n  want: %s", err.message, message)
		return
	}
}

func TestNewError2(t *testing.T) {
	code := codes.Internal
	reason := "Internal"
	message := "システム内部でエラーが発生しました。"
	trace := "trace"

	err, ok := ErrInternal.New(WithTrace(trace)).(*Error)
	if !ok {
		t.Errorf("Failed type assertion")
		return
	}

	if code != err.code {
		t.Errorf("\n  got: %s\n  want: %s", err.code, code)
		return
	}
	if reason != err.reason {
		t.Errorf("\n  got: %s\n  want: %s", err.reason, reason)
		return
	}
	if message != err.message {
		t.Errorf("\n  got: %s\n  want: %s", err.message, message)
		return
	}
}

func TestNewWrap1(t *testing.T) {
	code := codes.OK
	reason := ""
	message := ""

	i := ErrInternal.New(WithTrace("Internal"))
	w, ok := NewWrap(i, WithTrace("Wrap")).(*Error)
	if !ok {
		t.Errorf("Failed type assertion")
		return
	}

	if code != w.code {
		t.Errorf("\n  got: %s\n  want: %s", w.code, code)
		return
	}
	if reason != w.reason {
		t.Errorf("\n  got: %s\n  want: %s", w.reason, reason)
		return
	}
	if message != w.message {
		t.Errorf("\n  got: %s\n  want: %s", w.message, message)
		return
	}
}

// golang の標準 errors パッケージの Is 関数は, 第一引数で渡されるエラーはラップ対象を遡って比較する
// https://cs.opensource.google/go/go/+/master:src/errors/wrap.go;l=45-58
// これに従い, Wrap されているエラーでも正しく遡って比較されているかをテスト
func TestIs1(t *testing.T) {
	i1 := ErrInternal.New(WithTrace("Internal 1"))
	i2 := ErrInternal.New(WithTrace("Internal 2"))
	w1 := NewWrap(i2, WithTrace("Wrap"))
	w2 := NewWrap(i2, WithTrace("Wrap"))

	tests := []struct {
		want bool
		err1 error
		err2 error
	}{
		{want: true, err1: i1, err2: i2},
		{want: true, err1: w1, err2: w2},

		{want: true, err1: w1, err2: i1},
		{want: true, err1: w1, err2: i2},
		{want: true, err1: w2, err2: i1},
		{want: true, err1: w2, err2: i2},
		{want: false, err1: i1, err2: w1},
		{want: false, err1: i1, err2: w2},
		{want: false, err1: i2, err2: w1},
		{want: false, err1: i2, err2: w2},
	}
	for _, test := range tests {
		got := Is(test.err1, test.err2)
		if got != test.want {
			t.Errorf("[%v] == [%v]: got: %t, want: %t", test.err1, test.err2, got, test.want)
			return
		}
	}
}
