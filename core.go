package of

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/content"
)

type TypeHandleFunc = func(id string, data json.RawMessage) error

type Core interface {
	Context() context.Context
	Send(ctx context.Context, message *Message) error
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
}
