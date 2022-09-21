package content

//Message ...
//@Description:
type Message struct {
	Last    int64  `json:"last,omitempty"`
	Index   int64  `json:"index,omitempty"`
	Version int    `json:"version,omitempty"` // current info version
	Length  int    `json:"length,omitempty"`
	Data    []byte `json:"data,omitempty"`
}

func (m *Message) IsEmpty() bool {
	return m == nil || m.Length == 0
}

func (m *Message) SetIndex(index int64) *Message {
	m.Index = index
	return m
}

func (m *Message) SetVersion(version int) *Message {
	m.Version = version
	return m
}

func (m *Message) SetDataLength(data []byte) *Message {
	m.Data = data
	m.Length = len(data)
	return m
}

func (m *Message) SetLast(last int64) *Message {
	m.Last = last
	return m
}

func (m *Message) Revise() *Message {
	if m.IsEmpty() {
		return m
	}
	if (m.Length == 0 && len(m.Data) == m.Length) &&
		m.Index == 0 && m.Last == 0 {
		return nil
	}
	return m
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
