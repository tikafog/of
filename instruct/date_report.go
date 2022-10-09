package instruct

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/instruct"
)

const CurrentReportVersion = 1

// ReportType ...
// ENUM(none,state,message,max)
type ReportType uint32

type ReportData struct {
	Version int        `json:"version,omitempty"`
	Type    ReportType `json:"type,omitempty"`
	Last    int64      `json:"last,omitempty"`
}

func NewReportData(p ReportType) *ReportData {
	return &ReportData{
		Version: CurrentReportVersion,
		Type:    p,
	}
}

func (d ReportData) InstructType() Type {
	return instruct.TypeReport
}

func (d *ReportData) JSON() []byte {
	v, _ := json.Marshal(d)
	return v
}

var _ Data = (*ReportData)(nil)
