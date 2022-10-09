package instruct

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/buffers/instruct"
)

const CurrentCorrectVersion = 1

type CorrectData struct {
	Version   int             `json:"version,omitempty"`
	Type      content.Type    `json:"type,omitempty"`
	List      json.RawMessage `json:"list,omitempty"`
	StartUnix int64           `json:"start_unix,omitempty"`
	EndUnix   int64           `json:"end_unix,omitempty"`
}

func NewCorrectData(p content.Type, list json.RawMessage) *CorrectData {
	return &CorrectData{
		Version: CurrentCorrectVersion,
	}
}

func (d CorrectData) InstructType() Type {
	return instruct.TypeCorrect
}

func (d *CorrectData) SetEndUnix(endUnix int64) {
	d.EndUnix = endUnix
}

func (d *CorrectData) SetStartUnix(startUnix int64) {
	d.StartUnix = startUnix
}

var _ Data = (*CorrectData)(nil)
