package instruct

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
)

const CurrentCorrectVersion = 1

type CorrectData struct {
	Version   int             `json:"version,omitempty"`
	Type      content.Type    `json:"type,omitempty"`
	List      json.RawMessage `json:"list,omitempty"`
	StartUnix int64           `json:"start_unix,omitempty"`
	EndUnix   int64           `json:"end_unix,omitempty"`
}

func (c *CorrectData) SetEndUnix(endUnix int64) {
	c.EndUnix = endUnix
}

func (c *CorrectData) SetStartUnix(startUnix int64) {
	c.StartUnix = startUnix
}

func NewCorrectData(p content.Type, list json.RawMessage) *CorrectData {
	return &CorrectData{
		Version: CurrentCorrectVersion,
	}
}
