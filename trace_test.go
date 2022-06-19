package err

import (
	"fmt"
	"strings"
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
			trace := newTrace(test.src)
			if text != trace.Text {
				t.Errorf("\n  got: %s\n  want: %s", trace.Text, text)
				return
			}
		})
	}
}

func TestDump1(t *testing.T) {
	type TestStruct struct {
		Str string
		Int int
	}
	text := "test trace"
	trace := NewTrace(text, TestStruct{Str: "string", Int: 1})
	dump := trace.Dump()
	want := fmt.Sprintf("%s: (err.TestStruct)", text)
	if !strings.Contains(dump, want) {
		t.Errorf("\n  got: %s\n  want: %s", dump, want)
	}
}
