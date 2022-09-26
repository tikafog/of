package content

import (
	"encoding/json"
)

const MessageV2Version = 2

//MessageV2 ...
//@Description:
type MessageV2 struct {
	Last    int64           `json:"last,omitempty"`
	Index   int64           `json:"index,omitempty"`
	Version int             `json:"version,omitempty"` // current info version
	Length  int             `json:"length,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (m *MessageV2) IsEmpty() bool {
	return m == nil || m.Length == 0
}

func (m *MessageV2) SetIndex(index int64) *MessageV2 {
	m.Index = index
	return m
}

func (m *MessageV2) SetVersion(version int) *MessageV2 {
	m.Version = version
	return m
}

func (m *MessageV2) SetDataLength(data json.RawMessage) *MessageV2 {
	m.Data = data
	m.Length = len(data)
	return m
}

func (m *MessageV2) SetLast(last int64) *MessageV2 {
	m.Last = last
	return m
}

func (m *MessageV2) v1() *MessageV1 {
	v1 := new(MessageV1)
	v1.Last = m.Last
	v1.Index = m.Index
	v1.Version = MessageV1Version
	v1.Data = m.Data
	v1.Length = len(v1.Data)
	return v1
}

func (m *MessageV2) Revise() *MessageV2 {
	if m.IsEmpty() {
		return m
	}
	if (m.Length == 0 && len(m.Data) == m.Length) &&
		m.Index == 0 && m.Last == 0 {
		return nil
	}
	return m
}

func (m *MessageV2) current() *Message {
	return messageV2ToCurrentMessage(m)
}

// NewContentMessageV2
// @param []byte
// @return *MessageV2
//Decrypted use NewMessageV2 instead
func NewContentMessageV2(data json.RawMessage) *MessageV2 {
	msg := MessageV2{
		Version: MessageV2Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewMessageV2
// @param []byte
// @return *MessageV2
func NewMessageV2(data json.RawMessage) *MessageV2 {
	msg := MessageV2{
		Version: MessageV2Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV2WithDetail
// @param []byte
// @param int64
// @param int64
// @return *MessageV2
func NewContentMessageV2WithDetail(data json.RawMessage, index int64, last int64) *MessageV2 {
	msg := MessageV2{
		Last:    last,
		Version: MessageV2Version,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageV2AndLast
// @param []byte
// @param int64
// @return *MessageV2
func NewContentMessageV2AndLast(data json.RawMessage, last int64) *MessageV2 {
	msg := MessageV2{
		Version: MessageV2Version,
		Length:  len(data),
		Data:    data,
		Last:    last,
	}
	return &msg
}

// NewContentMessageV2Last
// @param int64
// @return *MessageV2
func NewContentMessageV2Last(last int64) *MessageV2 {
	msg := MessageV2{
		Version: MessageV2Version,
		Last:    last,
	}
	return &msg
}

// ParseMessageV2Data ...
// @param json.RawMessage
// @return []*T
// @return error
func ParseMessageV2Data(raw json.RawMessage) ([]byte, error) {
	t := new([]byte)
	err := json.Unmarshal(raw, t)
	return *t, err
}
