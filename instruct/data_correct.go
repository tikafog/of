package instruct

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
)

type CorrectData struct {
	DataVersion
	Type      content.Type    `json:"type,omitempty"`
	List      json.RawMessage `json:"list,omitempty"`
	StartUnix int64           `json:"start_unix,omitempty"`
	EndUnix   int64           `json:"end_unix,omitempty"`
}
