package error

type Option interface {
	Apply(*Options)
}

func NewOptions(sources ...Option) *Options {
	opts := &Options{sources: sources}
	for _, source := range sources {
		if source == nil {
			continue
		}
		source.Apply(opts)
	}
	return opts
}

type Options struct {
	Trace *Trace

	sources []Option
}

// Trace

func WithTrace(v interface{}) trace {
	t := newTrace(v)
	return trace{&t}
}

type trace struct {
	*Trace
}

func (o trace) Apply(opts *Options) {
	opts.Trace = o.Trace
}
