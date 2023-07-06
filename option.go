package ers

type WrapOption func(o *wrapOptions)

type wrapOptions struct {
	Trace any
}

// WithTrace sets the trace option.
func WithTrace(v any) WrapOption {
	return func(o *wrapOptions) {
		o.Trace = v
	}
}
