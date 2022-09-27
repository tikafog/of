package boot

import (
	"context"

	"github.com/tikafog/of"
	"github.com/tikafog/of/merr"
	"github.com/tikafog/of/option"
)

type emptyModule struct {
	name of.Name
}

var ErrEmptyModule = merr.New("empty module")

func NewEmptyModule(name of.Name) of.Module {
	return newEmptyLoader(name)
}

func newEmptyLoader(name of.Name) Loader {
	return &emptyModule{
		name: name,
	}
}

func (m emptyModule) RegisterEvent(event of.Event) error {
	return nil
}

func (m emptyModule) IsValid() bool {
	return false
}

func (m emptyModule) WithInit(o option.InitializeOption) of.ModuleStarter {
	return m
}

func (m emptyModule) WithOption(o option.Option) of.ModuleStarter {
	return m
}

func (m emptyModule) RegisterAPI(api of.API) error {
	return nil
}

func (m emptyModule) Init() error {
	return nil
}

func (m emptyModule) Destroy() {

}

func (m emptyModule) PreloadCore(core of.Core) error {
	return nil
}

func (m emptyModule) Run(ctx context.Context) error {
	return nil
}

func (m emptyModule) Name() of.Name {
	return m.name
}
