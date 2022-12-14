package eload

import (
	"context"
	"github.com/tikafog/of"
)

type Runner interface {
	Init(name of.Name) error
	InitAll() error
	StartRun(ctx context.Context, name of.Name) error
	StartAll(ctx context.Context) error
	Stop(name of.Name) error
	StopAll() error
}
