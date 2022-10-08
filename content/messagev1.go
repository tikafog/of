package content

import (
	"encoding/json"
)

const MessageV1Version = 1

// MessageV1 ...
// @Description:
type MessageV1 struct {
	Last    int64  `json:"last,omitempty"`
	Index   int64  `json:"index,omitempty"`
	Version int    `json:"version,omitempty"` // current info version
	Length  int    `json:"length,omitempty"`
	Data    []byte `json:"data,omitempty"`
}

func (m *MessageV1) IsEmpty() bool {
	return m == nil || (m.Last == 0 && m.Index == 0 && m.Length == 0 && m.Version == 0)
}

func (m *MessageV1) SetVersion(version int) *MessageV1 {
	m.Version = version
	return m
}

func (m *MessageV1) SetDataLength(data []byte) *MessageV1 {
	m.Data = data
	m.Length = len(data)
	return m
}

func (m *MessageV1) SetLast(last int64) *MessageV1 {
	m.Last = last
	return m
}

func (m *MessageV1) Revise() *MessageV1 {
	if m.IsEmpty() {
		return m
	}
	if (m.Length == 0 && len(m.Data) == m.Length) &&
		m.Last == 0 {
		return nil
	}
	return m
}

func (m *MessageV1) JSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *MessageV1) v1() *MessageV1 {
	return m
}

func (m *MessageV1) v2() *MessageV2 {
	v2 := new(MessageV2)
	v2.Last = m.Last
	v2.Index = m.Index
	v2.Version = MessageV2Version
	v2.Data = m.Data
	v2.Length = len(v2.Data)
	return v2
}

func (m *MessageV1) current() *Message {
	return messageV1ToCurrentMessage(m)
}

// NewContentMessageV1
// @param []byte
// @return *MessageV1
// Decrypted use NewMessageV1 instead
func NewContentMessageV1(data []byte) *MessageV1 {
	msg := MessageV1{
		Version: MessageV1Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewMessageV1
// @param []byte
// @return *MessageV1
func NewMessageV1(data []byte) *MessageV1 {
	msg := MessageV1{
		Version: MessageV1Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV1WithDetail
// @param []byte
// @param int64
// @param int64
// @return *MessageV1
func NewContentMessageV1WithDetail(data []byte, index int64, last int64) *MessageV1 {
	msg := MessageV1{
		Last:    last,
		Version: MessageV1Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV1AndLast
// @param []byte
// @param int64
// @return *MessageV1
func NewContentMessageV1AndLast(data []byte, last int64) *MessageV1 {
	msg := MessageV1{
		Version: MessageV1Version,
		Length:  len(data),
		Data:    data,
		Last:    last,
	}
	return &msg
}

// NewContentMessageV1Last
// @param int64
// @return *MessageV1
func NewContentMessageV1Last(last int64) *MessageV1 {
	msg := MessageV1{
		Version: MessageV1Version,
		Last:    last,
	}
	return &msg
}

// ParseMessageV1Data ...
// @param json.RawMessage
// @return []*T
// @return error
func ParseMessageV1Data(raw json.RawMessage) ([]byte, error) {
	return raw, nil
}
