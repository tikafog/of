package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
)

// ExtPin ...
// @Description:
type ExtPin struct {
	Status string `json:"status,omitempty"`
	Step   string `json:"step,omitempty"`
	Rid    string `json:"rid,omitempty"`
	Error  string `json:"error,omitempty"`
}

// ExtType ...
// @Description:
// @receiver ContentReport
// @return content.ExtType
func (c ExtPin) ExtType() content.ExtType {
	return content.ExtTypePin
}

// MarshalData ...
// @Description:
// @receiver ContentReport
// @return data
func (c ExtPin) MarshalData() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalData ...
// @Description:
// @receiver *ExtPin
// @param []byte
// @return error
func (c *ExtPin) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, c)
}

var _ ExtConvertable = (*ExtPin)(nil)
