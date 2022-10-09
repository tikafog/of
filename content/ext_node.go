package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/utils"
)

// NodeState ...
// ENUM(offline,online,max)
type NodeState int

// NodeType ...
// ENUM(server,adapter,client,box,mobile,max)
type NodeType int

// ExtNode ...
// @Description:
type ExtNode struct {
	State    NodeState `json:"state,omitempty"`
	Type     NodeType  `json:"type,omitempty"`
	PID      string    `json:"pid,omitempty"`
	CPUID    string    `json:"cpuid,omitempty"`
	Addr     string    `json:"addr,omitempty"`
	Addrs    []string  `json:"addrs,omitempty"`
	Length   int       `json:"length,omitempty"`
	Data     []byte    `json:"data,omitempty"`
	NodeType string    `json:"node_type,omitempty"`
	TimeUnix int64     `json:"time_unix,omitempty"`
}

// ExtType ...
// @Description:
// @receiver ExtNode
// @return content.ExtType
func (c *ExtNode) ExtType() content.ExtType {
	return content.ExtTypeNode
}

// JSON
// @receiver *ExtNode
// @return []byte
func (c *ExtNode) JSON() []byte {
	return utils.Must(json.Marshal(c))
}

// Struct
// @receiver *ExtNode
// @param []byte
// @return error
func (c *ExtNode) Struct(data []byte) error {
	return json.Unmarshal(data, c)
}

// MarshalData ...
// @Description:
// @receiver ExtNode
// @return data
func (c *ExtNode) MarshalData() (data []byte, err error) {
	return json.Marshal(c)
}

// UnmarshalData ...
// @Description:
// @receiver *ExtNode
// @param []byte
// @return error
func (c *ExtNode) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, c)
}

var _ ExtConverter = (*ExtNode)(nil)
