package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/utils"
)

// ExtResource ...
// @Description:
type ExtResource struct {
	RID      string `json:"rid,omitempty"`
	Status   string `json:"status,omitempty"`
	Step     string `json:"step,omitempty"`
	Error    string `json:"error,omitempty"`
	TimeUnix int64  `json:"time_unix,omitempty"`
}

// ExtType ...
// @Description:
// @receiver ContentReport
// @return content.ExtType
func (c *ExtResource) ExtType() content.ExtType {
	return content.ExtTypeResource
}

// JSON
// @receiver *ExtResource
// @return []byte
func (c *ExtResource) JSON() []byte {
	return utils.Must(json.Marshal(c))
}

// Struct
// @receiver *ExtResource
// @param []byte
// @return error
func (c *ExtResource) Struct(data []byte) error {
	return json.Unmarshal(data, c)
}

// MarshalData ...
// @Description:
// @receiver ContentReport
// @return data
func (c *ExtResource) MarshalData() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalData ...
// @Description:
// @receiver *ExtResource
// @param []byte
// @return error
func (c *ExtResource) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, c)
}

var _ ExtConverter = (*ExtResource)(nil)
