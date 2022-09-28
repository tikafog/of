package of

import (
	"context"
)

type BootstrapLsOption func(*BootstrapLsSetting)

type BootstrapLsSetting struct {
	Limit int
}

// SetBootstrapLimit ...
// @param int
// @return BootstrapLsOption
func SetBootstrapLimit(limit int) BootstrapLsOption {
	return func(option *BootstrapLsSetting) {
		option.Limit = limit
	}
}

type Bootstrap interface {
	Add(ctx context.Context, addrs ...string) error
	Ls(ctx context.Context, opts ...BootstrapLsOption) (int, []string, error)
	//Range(ctx context.Context, fn func(addr string) bool)
	Remove(ctx context.Context, addr ...string) error
}
