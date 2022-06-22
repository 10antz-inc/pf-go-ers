package option

import (
	"fmt"

	"github.com/tys-muta/go-opt"
)

type trace struct {
	Trace any
}

var _ opt.Option = (*trace)(nil)

func WithTrace(v any) trace {
	return trace{v}
}

func (o trace) Validate() error {
	if o.Trace == nil {
		return fmt.Errorf("trace is undefined")
	}
	return nil
}

func (o trace) Apply(options any) {
	switch v := options.(type) {
	case *WrapOptions:
		v.Trace = o.Trace
	}
}
