package ers

import (
	"testing"
)

func TestNewTrace1(t *testing.T) {
	text := "text"
	trace := NewTrace(text)

	if text != trace.Text {
		t.Errorf("\n  got: %s\n  want: %s", trace.Text, text)
		return
	}
}

func TestNewTrace2(t *testing.T) {
	text := "test trace"
	trace := NewTrace(text)

	type TestString string

	tests := []struct {
		src interface{}
	}{
		{src: text},
		{src: trace},
		{src: *trace},
		{src: TestString(text)},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			trace := NewTrace(test.src)
			if text != trace.Text {
				t.Errorf("\n  got: %s\n  want: %s", trace.Text, text)
				return
			}
		})
	}
}
