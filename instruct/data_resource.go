package instruct

import (
	"github.com/tikafog/of/buffers/instruct"
)

const CurrentResourceVersion = 1

type ResourceData struct {
	Version  int      `json:"version,omitempty"`
	List     []string `json:"list,omitempty"`
	Step     uint32   `json:"step,omitempty"`
	Priority int      `json:"priority,omitempty"`
	Relate   string   `json:"relate,omitempty"`
}

func NewResourceData() *ResourceData {
	return &ResourceData{
		Version: CurrentResourceVersion,
	}
}

func (d ResourceData) InstructType() Type {
	return instruct.TypeResource
}

var _ Data = (*ResourceData)(nil)
