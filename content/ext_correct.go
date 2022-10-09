package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/utils"
)

// ExtCorrect ...
// @Description:
type ExtCorrect struct {
	Type      content.Type `json:"type,omitempty"`
	StartUnix int64        `json:"start_unix,omitempty"`
	EndUnix   int64        `json:"end_unix,omitempty"`
	Hash      []byte       `json:"hash,omitempty"`
}

// JSON
// @receiver *ExtCorrect
// @return []byte
func (c *ExtCorrect) JSON() []byte {
	return utils.Must(json.Marshal(c))
}

// Struct
// @receiver *ExtCorrect
// @param []byte
// @return error
func (c *ExtCorrect) Struct(data []byte) error {
	return json.Unmarshal(data, c)
}

// ExtType ...
// @Description:
// @receiver ExtCorrect
// @return content.ExtType
func (c *ExtCorrect) ExtType() content.ExtType {
	return content.ExtTypeCorrect
}

// MarshalData ...
// @Description:
// @receiver ExtCorrect
// @return data
func (c *ExtCorrect) MarshalData() (data []byte, err error) {
	return json.Marshal(c)
}

// UnmarshalData ...
// @Description:
// @receiver *ExtCorrect
// @param []byte
// @return error
func (c *ExtCorrect) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, c)
}

var _ ExtConverter = (*ExtCorrect)(nil)
