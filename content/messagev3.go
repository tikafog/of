package content

import (
	"encoding/json"

	"github.com/tikafog/of/utils"
)

const MessageV3Version = 3

//MessageV3 ...
//@Description:
type MessageV3[T any] struct {
	Last  int64 `json:"last,omitempty"`
	Index int64 `json:"index,omitempty"`
	Data  []*T  `json:"data,omitempty"`
}

func (m *MessageV3[T]) IsEmpty() bool {
	return m == nil || len(m.Data) == 0
}

func (m *MessageV3[T]) SetIndex(index int64) *MessageV3[T] {
	m.Index = index
	return m
}

func (m *MessageV3[T]) SetLast(last int64) *MessageV3[T] {
	m.Last = last
	return m
}

func (m *MessageV3[T]) Revise() *MessageV3[T] {
	if m.IsEmpty() {
		return m
	}
	if (len(m.Data) == 0) &&
		m.Index == 0 && m.Last == 0 {
		return nil
	}
	return m
}

func (m *MessageV3[T]) v1() *MessageV1 {
	v1 := new(MessageV1)
	v1.Index = m.Index
	v1.Last = m.Last
	v1.Version = MessageV1Version
	v1.Data = utils.Must(json.Marshal(m.Data))
	v1.Length = len(v1.Data)
	return v1
}

func (m *MessageV3[T]) v2() *MessageV2 {
	v2 := new(MessageV2)
	v2.Index = m.Index
	v2.Last = m.Last
	v2.Version = MessageV2Version
	v2.Data = utils.Must(json.Marshal(m.Data))
	v2.Length = len(v2.Data)
	return v2
}

// NewContentMessageV3
// @param []byte
// @return *MessageV3[T]
//Decrypted use NewMessageV3[T] instead
func NewContentMessageV3[T any](data []*T) *MessageV3[T] {
	msg := MessageV3[T]{
		Data: data,
	}
	return &msg
}

// NewMessageV3
// @param []byte
// @return *MessageV3[T]
func NewMessageV3[T any](data []*T) *MessageV3[T] {
	msg := MessageV3[T]{
		Data: data,
	}
	return &msg
}

// NewContentMessageV3WithDetail
// @param []byte
// @param int64
// @param int64
// @return *MessageV3[T]
func NewContentMessageV3WithDetail[T any](data []*T, index int64, last int64) *MessageV3[T] {
	msg := MessageV3[T]{
		Last:  last,
		Index: index,
		Data:  data,
	}
	return &msg
}

// NewContentMessageV3AndLast
// @param []byte
// @param int64
// @return *MessageV3[T]
func NewContentMessageV3AndLast[T any](data []*T, last int64) *MessageV3[T] {
	msg := MessageV3[T]{
		Data: data,
		Last: last,
	}
	return &msg
}

// NewContentMessageV3Last
// @param int64
// @return *MessageV3[T]
func NewContentMessageV3Last[T any](last int64) *MessageV3[T] {
	msg := MessageV3[T]{
		Last: last,
	}
	return &msg
}
