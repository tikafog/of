package content

import (
	"encoding/json"
)

const MessageV3Version = 3

//MessageV3 ...
//@Description:
type MessageV3 struct {
	Last    int64           `json:"last,omitempty"`
	Index   int64           `json:"index,omitempty"`
	Version int             `json:"version,omitempty"` // current info version
	Length  int             `json:"length,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (m *MessageV3) IsEmpty() bool {
	return m == nil || m.Length == 0
}

func (m *MessageV3) SetIndex(index int64) *MessageV3 {
	m.Index = index
	return m
}

func (m *MessageV3) SetVersion(version int) *MessageV3 {
	m.Version = version
	return m
}

func (m *MessageV3) SetDataLength(data []byte) *MessageV3 {
	m.Data = data
	m.Length = len(data)
	return m
}

func (m *MessageV3) SetLast(last int64) *MessageV3 {
	m.Last = last
	return m
}

func (m *MessageV3) Revise() *MessageV3 {
	if m.IsEmpty() {
		return m
	}
	if (m.Length == 0 && len(m.Data) == m.Length) &&
		m.Index == 0 && m.Last == 0 {
		return nil
	}
	return m
}

// NewContentMessageV3
// @param []byte
// @return *MessageV3
//Decrypted use NewMessageV3 instead
func NewContentMessageV3(data []byte) *MessageV3 {
	msg := MessageV3{
		Version: MessageV3Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewMessageV3
// @param []byte
// @return *MessageV3
func NewMessageV3(data []byte) *MessageV3 {
	msg := MessageV3{
		Version: MessageV3Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV3WithDetail
// @param []byte
// @param int64
// @param int64
// @return *MessageV3
func NewContentMessageV3WithDetail(data []byte, index int64, last int64) *MessageV3 {
	msg := MessageV3{
		Last:    last,
		Version: MessageV3Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV3AndLast
// @param []byte
// @param int64
// @return *MessageV3
func NewContentMessageV3AndLast(data []byte, last int64) *MessageV3 {
	msg := MessageV3{
		Version: MessageV3Version,
		Length:  len(data),
		Data:    data,
		Last:    last,
	}
	return &msg
}

// NewContentMessageV3Last
// @param int64
// @return *MessageV3
func NewContentMessageV3Last(last int64) *MessageV3 {
	msg := MessageV3{
		Version: MessageV3Version,
		Last:    last,
	}
	return &msg
}
