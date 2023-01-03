package eload

import (
	"encoding/json"
	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

type OptionLoader interface {
	WithOption(option.Option) of.ModuleStarter
}

type ConfigLoader interface {
	LoadConfig(message json.RawMessage) (json.RawMessage, error)
}

type Loader interface {
	OptionLoader
	ConfigLoader
	Injector
	NameRegister
	ModuleInterface
}

type NameRegister interface {
	Names() []of.Name
	Register(id uint64, name of.Name) error
}

type ModuleInterface interface {
	Add(id uint64, m of.Module) error
	Get(id uint64) (of.Module, bool)
	GetByName(name of.Name) (of.Module, bool)
}
