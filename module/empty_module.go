package module

import (
	"encoding/json"
)

type emptyModule struct {
	name Name
}

func (n emptyModule) Valid() bool {
	return false
}

func (n emptyModule) Start() error {
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
