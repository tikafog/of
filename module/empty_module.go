package module

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/module/option"
)

type emptyModule struct {
	name Name
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

func (n emptyModule) Name() Name {
	return n.name
}

func (n emptyModule) HandleData(id string, data json.RawMessage) error {
	return nil
}

func EmptyModule(name Name) Module {
	return &emptyModule{
		name: name,
	}
}
