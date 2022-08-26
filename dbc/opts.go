package dbc

const MinTimeoutSec = 60

type Option struct {
	timeout int64
	debug   bool
}

type Opts func(opt *Option)

func parseOption(opts []Opts) *Option {
	o := &Option{
		timeout: 300,
		debug:   false,
	}
	for i := range opts {
		opts[i](o)
	}
	return o
}

func Debug(b bool) Opts {
	return func(opt *Option) {
		opt.debug = b
	}
}

func TimeoutOpt(t int64) Opts {
	return func(opt *Option) {
		opt.timeout = t
	}
}
