package module

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of"
	"github.com/tikalink/of/option"
)

type emptyModule struct {
	name of.Name
}

func (n emptyModule) Option(op option.Option) error {
	return nil
}

func (n emptyModule) Valid() bool {
	return false
}

func (n emptyModule) Run(ctx context.Context) error {
	return nil
}

func (n emptyModule) Name() of.Name {
	return n.name
}

func (n emptyModule) HandleData(id string, data json.RawMessage) error {
	return nil
}

func EmptyModule(name of.Name) of.Module {
	return &emptyModule{
		name: name,
	}
}
