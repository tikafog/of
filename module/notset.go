package module

import (
	"github.com/tikafog/of"
)

type notsetModule struct{}

func (n notsetModule) WaitEvent(name of.Name, args ...of.Arg) error {
	return nil
}

func (n notsetModule) IsNil() bool {
	return false
}

func (n notsetModule) IsRunning() bool {
	return false
}

func (n notsetModule) Name() of.Name {
	return of.NameNotSet
}

var NotSet of.Module = &notsetModule{}
