package dbc

const MinTimeoutSec = 60

type Option struct {
	timeout int64
}

type Opts func(opt *Option)

func opt(opts []Opts) *Option {
	o := &Option{
		timeout: 300,
	}
	for i := range opts {
		opts[i](o)
	}
	return o
}

func TimeoutOpt(t int64) Opts {
	return func(opt *Option) {
		opt.timeout = t
	}
}
