package instruct

import (
	"github.com/tikafog/of/buffers/instruct"
)

const CurrentResourceVersion = 1

type ResourceData struct {
	Version int      `json:"version,omitempty"`
	List    []string `json:"list,omitempty"`
	Action  int      `json:"action,omitempty"`
	Last    int64    `json:"last,omitempty"`
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
