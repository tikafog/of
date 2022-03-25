package module

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

type emptyModule struct {
	name of.Name
}

var ErrEmptyModule = errors.New("empty module")

func (m emptyModule) Query(limit int, last int64) ([]byte, error) {
	return nil, ErrEmptyModule
}

func (m emptyModule) WaitEvent(name of.Name, args ...of.Arg) error {
	return ErrEmptyModule
}

func (m emptyModule) Data(limit int, last int64) ([]byte, error) {
	return nil, ErrEmptyModule
}

func (m emptyModule) IsRunning() bool {
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

func (m emptyModule) SetCore(core of.Core) error {
	return nil
}

func (m emptyModule) IsNil() bool {
	return false
}

func (m emptyModule) Run(ctx context.Context) error {
	return nil
}

func (m emptyModule) Name() of.Name {
	return m.name
}

func (m emptyModule) HandleData(id string, data json.RawMessage) error {
	return nil
}

func NewEmptyModule(name of.Name) of.Module {
	return newEmptyLoader(name)
}

func newEmptyLoader(name of.Name) Loader {
	return &emptyModule{
		name: name,
	}
}
