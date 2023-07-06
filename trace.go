package ers

import (
	"fmt"
)

var (
	// T 関数は, NewTrace 関数のエイリアス.
	T = NewTrace
)

type Trace struct {
	Text   string
	Values []any
}

func NewTrace(src any) *Trace {
	switch v := src.(type) {
	case string:
		return &Trace{Text: v}
	case []byte:
		return &Trace{Text: string(v)}
	case error:
		return &Trace{Text: v.Error()}
	case *Trace:
		if v != nil {
			return &Trace{Text: v.Text, Values: v.Values}
		}
	case Trace:
		return &v
	}
	return &Trace{Text: fmt.Sprintf("%s", src)}
}
