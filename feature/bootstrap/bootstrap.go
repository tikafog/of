package bootstrap

import (
	"context"

	"github.com/tikafog/of"
)

// Level ...
// ENUM(core,speed,normal,max)
type Level uint32

type LsOption func(*LsSetting)

type LsSetting struct {
	Index int64
	Limit int
}

// LsLimit ...
// @param int
// @return BootstrapLsOption
func LsLimit(limit int) LsOption {
	return func(option *LsSetting) {
		option.Limit = limit
	}
}

// LsIndex ...
// @param int
// @return BootstrapLsOption
func LsIndex(index int64) LsOption {
	return func(option *LsSetting) {
		option.Index = index
	}
}

type UpgradeOption func(*UpgradeSetting)

type UpgradeSetting struct {
	Level Level
	Addrs []string
}

// UpgradeLevel ...
// @param bootstrap.Level
// @return BootstrapUpgradeOption
func UpgradeLevel(level Level) UpgradeOption {
	return func(option *UpgradeSetting) {
		option.Level = level
	}
}

// UpgradeAddr ...
// @param ...string
// @return BootstrapUpgradeOption
func UpgradeAddr(addrs ...string) UpgradeOption {
	return func(option *UpgradeSetting) {
		option.Addrs = addrs
	}
}

// Bootstrap is a bootstrap interface
type Bootstrap interface {
	Add(ctx context.Context, addrs ...string) error
	Ls(ctx context.Context, opts ...LsOption) (int64, []string, error)
	Upgrade(ctx context.Context, id of.ID, opts ...UpgradeOption) error
	Remove(ctx context.Context, addr ...string) error
}
