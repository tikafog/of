package module

import (
	"github.com/tikafog/of"
)

type notsetModule struct{}

func (n notsetModule) Valid() bool {
	return false
}

func (n notsetModule) IsRunning() bool {
	return false
}

func (n notsetModule) Name() of.Name {
	return of.NameNotSet
}

var NotSet of.ModuleStater = &notsetModule{}
