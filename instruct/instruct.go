package instruct

import (
	"encoding/json"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tikafog/of/buffers/instruct"
	"github.com/tikafog/of/utils"
	"github.com/tikafog/of/version"
)

// Type ...
type Type = instruct.Type

type metaParser interface {
	parseMetaInstruct(m *metaInstruct) error
}

type metaInstruct struct {
	Version string          `json:"version,omitempty"`
	To      string          `json:"to,omitempty"`
	Length  int             `json:"length,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Type    Type            `json:"type,omitempty"`
	Last    int64           `json:"last,omitempty"`
}

type Data interface {
	InstructType() Type
}

type Instructor interface {
	GetTo() string
	GetType() instruct.Type
	GetData() any
	GetLast() int64
	JSON() []byte
	Bytes() []byte
}

type DataInstructor interface {
	CorrectData | ResourceData | ReportData
	Data
}

// Instruct ...
type Instruct[T DataInstructor] struct {
	meta *metaInstruct
	To   string `json:"to,omitempty"`
	Type Type   `json:"type,omitempty"`
	Data *T     `json:"data,omitempty"`
	Last int64  `json:"last,omitempty"`
}

func (i *Instruct[T]) GetTo() string {
	return i.To
}

func (i *Instruct[T]) GetType() Type {
	return i.Type
}

func (i *Instruct[T]) GetData() any {
	return i.Data
}

func (i *Instruct[T]) GetLast() int64 {
	return i.Last
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (t *metaInstruct) Bytes() []byte {
	t.Version = version.VersionTwo
	return instructToBytes(t)
}

// SetVersion ...
// @receiver *metaInstruct
// @param string
// @return *metaInstruct
func (t *metaInstruct) SetVersion(version string) *metaInstruct {
	t.Version = version
	return t
}

// JSON ...
// @receiver *metaInstruct
// @return []byte
// @return error
func (t *metaInstruct) JSON() ([]byte, error) {
	t.Version = version.VersionOne
	return json.Marshal(t)
}

// SetTo ...
// @receiver *Instruct[T]
// @param string
func (i *Instruct[T]) SetTo(To string) {
	i.To = To
}

// MarshalData ...
// @Description:
// @receiver ContentReport
// @return data
func (i *Instruct[T]) MarshalData() ([]byte, error) {
	return json.Marshal(i)
}

// UnmarshalData ...
// @Description:
// @receiver *Instruct
// @param []byte
// @return error
func (i *Instruct[T]) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, i)
}

// SetData
// @receiver *Instruct[T]
// @param []byte
// @return *Instruct[T]
func (i *Instruct[T]) SetData(data *T) *Instruct[T] {
	i.Data = data
	return i
}

// SetType
// @receiver *Instruct[T]
// @param instruct.Type
// @return *Instruct[T]
func (i *Instruct[T]) SetType(p instruct.Type) *Instruct[T] {
	i.Type = p
	//i.meta.Type = i.Type
	return i
}

// JSON ...
// @Description:
// @receiver Instruct[T]
// @return []byte
// @return error
func (i *Instruct[T]) JSON() []byte {
	return utils.Must(i.metaInstruct().JSON())
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (i *Instruct[T]) Bytes() []byte {
	return i.metaInstruct().Bytes()
}

func (i *Instruct[T]) parseMetaInstruct(m *metaInstruct) error {
	i.meta = m
	i.Type = m.Type
	i.To = m.To
	i.Last = m.Last
	if m.Length != 0 && i.Data == nil {
		i.Data = new(T)
		return json.Unmarshal(m.Data, i.Data)
	}
	return nil
}

func (i *Instruct[T]) metaInstruct() *metaInstruct {
	if i.meta == nil {
		i.meta = new(metaInstruct)
		i.meta.Type = i.Type
		i.meta.To = i.To
		i.meta.Data = utils.Must(json.Marshal(i.Data))
		i.meta.Length = len(i.meta.Data)
	}
	return i.meta
}

func instructToBytes(c *metaInstruct) []byte {
	builder := flatbuffers.NewBuilder(0)
	_data := builder.CreateByteString(c.Data)
	_version := builder.CreateString(c.Version)
	_to := builder.CreateString(c.To)
	instruct.InstructStart(builder)
	instruct.InstructAddData(builder, _data)
	instruct.InstructAddTo(builder, _to)
	instruct.InstructAddVersion(builder, _version)
	instruct.InstructAddType(builder, c.Type)
	instruct.InstructAddLast(builder, c.Last)
	builder.Finish(instruct.InstructEnd(builder))
	return builder.FinishedBytes()
}

// NewInstruct ...
// @return *Instruct[T]
func NewInstruct[T DataInstructor]() *Instruct[T] {
	inst := new(Instruct[T])
	inst.Type = Data.InstructType(*new(T))
	return inst
}
