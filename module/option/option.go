package option

type Option interface {
	Repo() string
}

func Repo(repo string) func(o Option) Option {
	return func(o Option) Option {
		op := o.(option)
		op.repo = repo
		return op
	}
}

func Default() Option {
	return &option{}
}

type option struct {
	repo string
}

func (o option) Repo() string {
	return o.repo
}
