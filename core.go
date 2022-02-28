package of

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/content"
)

type TypeHandleFunc = func(id string, data json.RawMessage) error

type Core interface {
	Connection
	Context() context.Context
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
}
