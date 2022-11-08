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

type DataSource interface {
	CorrectData | ResourceData | ReportData
}

type dataAble interface {
	InstructType() Type
}

type Getter interface {
	To() string
	Type() instruct.Type
	Data() json.RawMessage
	Last() int64
	Version() string
	JSON() []byte
	Bytes() []byte
}

type Setter interface {
	SetTo(string)
	SetType(instruct.Type)
	SetData(message json.RawMessage)
	SetLast(int64)
	SetVersion(string)
}

type Instructor interface {
	Getter
	Setter
}

// Instruct ...
type Instruct[T DataSource] struct {
	meta *metaInstruct
	//To   string `json:"to,omitempty"`
	//Type Type   `json:"type,omitempty"`
	dataSource *T
	//Last int64  `json:"last,omitempty"`
}

func (i *Instruct[T]) Version() string {
	return i.meta.Version
}

func (i *Instruct[T]) SetVersion(v string) {
	i.meta.Version = v
}

// SetLast ...
// @receiver *Instruct[T]
// @param int64
func (i *Instruct[T]) SetLast(last int64) {
	i.meta.Last = last
}

// To ...
// @receiver *Instruct[T]
// @return string
func (i *Instruct[T]) To() string {
	return i.meta.To
}

// Type ...
// @receiver *Instruct[T]
// @return Type
func (i *Instruct[T]) Type() Type {
	return i.meta.Type
}

// Data ...
// @receiver *Instruct[T]
// @return json.RawMessage
func (i *Instruct[T]) Data() json.RawMessage {
	return i.meta.Data
}

// Last ...
// @receiver *Instruct[T]
// @return int64
func (i *Instruct[T]) Last() int64 {
	return i.meta.Last
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
	i.meta.To = To
}

// MarshalData ...
// @Description:
// @receiver ContentReport
// @return data
func (i *Instruct[T]) MarshalData() ([]byte, error) {
	return json.Marshal(i.meta)
}

// UnmarshalData ...
// @Description:
// @receiver *Instruct
// @param []byte
// @return error
func (i *Instruct[T]) UnmarshalData(data []byte) error {
	i.meta = new(metaInstruct)
	return json.Unmarshal(data, i.meta)
}

// SetData ...
// @receiver *Instruct[T]
// @param json.RawMessage
func (i *Instruct[T]) SetData(data json.RawMessage) {
	i.meta.Data = data
	i.meta.Length = len(data)
}

// SetDataSource
// @receiver *Instruct[T]
// @param []byte
// @return *Instruct[T]
func (i *Instruct[T]) SetDataSource(data *T) *Instruct[T] {
	i.dataSource = data
	i.SetData(utils.Must(json.Marshal(data)))
	return i
}

// SetType ...
// @receiver *Instruct[T]
// @param instruct.Type
func (i *Instruct[T]) SetType(p instruct.Type) {
	i.meta.Type = p
}

// JSON ...
// @Description:
// @receiver Instruct[T]
// @return []byte
// @return error
func (i *Instruct[T]) JSON() []byte {
	return utils.Must(i.meta.JSON())
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (i *Instruct[T]) Bytes() []byte {
	return i.meta.Bytes()
}

func (i *Instruct[T]) parseMetaInstruct(m *metaInstruct) error {
	i.meta = m
	if m.Length != 0 && m.Data == nil {
		i.dataSource = new(T)
		return json.Unmarshal(m.Data, i.dataSource)
	}
	return nil
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
func NewInstruct[T DataSource]() *Instruct[T] {
	inst := new(Instruct[T])
	inst.meta = new(metaInstruct)
	inst.SetType(dataType(new(T)))
	return inst
}

func CastInstruct[T DataSource](instructor Instructor) (*Instruct[T], bool) {
	v, ok := any(instructor).(*Instruct[T])
	return v, ok
}

func dataType(a any) instruct.Type {
	p, ok := a.(dataAble)
	if ok {
		return dataAble.InstructType(p)
	}
	return instruct.TypeMax
}

var _ Instructor = (*Instruct[ReportData])(nil)
var _ Instructor = (*Instruct[CorrectData])(nil)
var _ Instructor = (*Instruct[ResourceData])(nil)
