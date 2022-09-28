package instruct

import (
	"encoding/json"
)

//ReportType ...
//ENUM(none,state,message,max)
type ReportType uint32

type ReportData struct {
	VersionData
	Type ReportType `json:"type,omitempty"`
	Last int64      `json:"last,omitempty"`
}

func (d ReportData) JSON() []byte {
	v, _ := json.Marshal(d)
	return v
}
