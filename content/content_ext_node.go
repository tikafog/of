package content

import (
	"encoding/json"

	"github.com/tikalink/of/buffers/content"
)

// NodeState ...
// ENUM(initialized,start,online,offline,max)
type NodeState int

// NodeType ...
// ENUM(server,adapter,box,mobile,max)
type NodeType int

//ExtNode ...
//@Description:
type ExtNode struct {
	State    NodeState       `json:"state,omitempty"`
	Type     NodeType        `json:"type,omitempty"`
	PID      string          `json:"pid,omitempty"`
	CPUID    string          `json:"cpuid,omitempty"`
	Addr     string          `json:"addr,omitempty"`
	Addrs    []string        `json:"addrs,omitempty"`
	Length   int             `json:"length,omitempty"`
	Data     json.RawMessage `json:"data,omitempty"`
	NodeType string          `json:"node_type,omitempty"`
}

// ExtType ...
// @Description:
// @receiver ExtNode
// @return content.ExtType
func (c ExtNode) ExtType() content.ExtType {
	return content.ExtTypeNode
}

// MarshalData ...
// @Description:
// @receiver ExtNode
// @return data
func (c ExtNode) MarshalData() (data []byte, err error) {
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

var _ ExtConvertable = (*ExtNode)(nil)
