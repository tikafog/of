package of

import (
	"context"

	"github.com/tikafog/of/dbc/bootnode/bootstrap"
)

type BootstrapLsOption func(*BootstrapLsSetting)

type BootstrapLsSetting struct {
	Index int
	Limit int
}

// BootstrapLsLimit ...
// @param int
// @return BootstrapLsOption
func BootstrapLsLimit(limit int) BootstrapLsOption {
	return func(option *BootstrapLsSetting) {
		option.Limit = limit
	}
}

// BootstrapLsIndex ...
// @param int
// @return BootstrapLsOption
func BootstrapLsIndex(index int) BootstrapLsOption {
	return func(option *BootstrapLsSetting) {
		option.Index = index
	}
}

type BootstrapUpgradeOption func(*BootstrapUpgradeSetting)

type BootstrapUpgradeSetting struct {
	Level bootstrap.Level
	Addrs []string
}

// BootstrapUpgradeLevel ...
// @param bootstrap.Level
// @return BootstrapUpgradeOption
func BootstrapUpgradeLevel(level bootstrap.Level) BootstrapUpgradeOption {
	return func(option *BootstrapUpgradeSetting) {
		option.Level = level
	}
}

// BootstrapUpgradeAddr ...
// @param ...string
// @return BootstrapUpgradeOption
func BootstrapUpgradeAddr(addrs ...string) BootstrapUpgradeOption {
	return func(option *BootstrapUpgradeSetting) {
		option.Addrs = addrs
	}
}

// Bootstrap is a bootstrap interface
type Bootstrap interface {
	Add(ctx context.Context, addrs ...string) error
	Ls(ctx context.Context, opts ...BootstrapLsOption) (int, []string, error)
	Upgrade(ctx context.Context, id ID, opts ...BootstrapUpgradeOption) error
	Remove(ctx context.Context, addr ...string) error
}
