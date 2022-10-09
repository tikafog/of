package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
)

// ExtCorrect ...
// @Description:
type ExtCorrect struct {
	Type      content.Type `json:"type,omitempty"`
	StartUnix int64        `json:"start_unix,omitempty"`
	EndUnix   int64        `json:"end_unix,omitempty"`
	Hash      []byte       `json:"hash,omitempty"`
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
