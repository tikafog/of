package eload

import (
	"context"
	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

// RunnerState ...
// ENUM(init,run,stop,max)
type RunnerState int

// RunnerHookFunc ...
type RunnerHookFunc = func(state RunnerState, module of.Module, args ...any) error

// Runner ...
type Runner interface {
	of.Module
	Init(option option.Option) error
	Run(ctx context.Context) error
	Stop()
}

type Runners interface {
	InitAll(option option.Option, fns ...RunnerHookFunc) error
	Init(name of.Name, option option.Option) error
	RunAll(ctx context.Context, fns ...RunnerHookFunc) error
	Run(ctx context.Context, name of.Name) error
	StopAll(fns ...RunnerHookFunc)
	Stop(name of.Name)
}
