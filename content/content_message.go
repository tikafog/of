package content

//Message ...
//@Description:
type Message struct {
	StartUnix int64  `json:"start_unix,omitempty"`
	Last      int64  `json:"last,omitempty"`
	Index     int64  `json:"index,omitempty"`
	Version   int    `json:"version,omitempty"` // current info version
	Length    int    `json:"length,omitempty"`
	Data      []byte `json:"data,omitempty"`
}

func (m *Message) IsEmpty() bool {
	return m == nil || m.Length == 0
}

func (m *Message) SetStartUnix(startUnix int64) *Message {
	m.StartUnix = startUnix
	return m
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

// EmptyMessage ...
var EmptyMessage = (*Message)(nil)