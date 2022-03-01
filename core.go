package of

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/content"
)

type TypeHandleFunc = func(id string, data json.RawMessage) error
type TypeEventFunc = func(r *EventRequest) error

type Core interface {
	Connection
	Context() context.Context
	//Ask(ctx context.Context, m *AskRequest) error
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
	RegisterEventHandler(from Name, fn TypeEventFunc) error
	Event(ctx context.Context, n Name, r *EventRequest) error
}
