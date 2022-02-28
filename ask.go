package of

import (
	"encoding/json"

	"github.com/tikalink/of/content"
)

type AskRequest struct {
	Type content.Type    `json:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}
