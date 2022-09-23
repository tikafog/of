package content

import (
	"encoding/json"
)

const CurrentDataVersion = MessageV2Version

type Message = MessageV2

//metaMessage ...
//@Description:
type metaMessage struct {
	Last    int64           `json:"last,omitempty"`
	Index   int64           `json:"index,omitempty"`
	Version int             `json:"version,omitempty"` // current info version
	Length  int             `json:"length,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// NewContentMessage
// @param []byte
// @return *Message
//Decrypted use NewMessage instead
func NewContentMessage(data []byte) *Message {
	msg := Message{
		Version: CurrentDataVersion,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewMessage
// @param []byte
// @return *Message
func NewMessage(data []byte) *Message {
	msg := Message{
		Version: CurrentDataVersion,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageWithDetail
// @param []byte
// @param int64
// @param int64
// @return *Message
func NewContentMessageWithDetail(data []byte, index int64, last int64) *Message {
	msg := Message{
		Last:    last,
		Index:   index,
		Version: CurrentDataVersion,
		Length:  len(data),
		Data:    data,
	}
	return &msg
}

// NewContentMessageAndLast
// @param []byte
// @param int64
// @return *Message
func NewContentMessageAndLast(data []byte, last int64) *Message {
	msg := Message{
		Version: CurrentDataVersion,
		Length:  len(data),
		Data:    data,
		Last:    last,
	}
	return &msg
}

// NewContentMessageLast
// @param int64
// @return *Message
func NewContentMessageLast(last int64) *Message {
	msg := Message{
		Version: CurrentDataVersion,
		Last:    last,
	}
	return &msg
}

// EmptyMessage ...
var EmptyMessage = (*Message)(nil)
