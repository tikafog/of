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
}

// Instruct ...
type Instruct[T any] struct {
	meta *metaInstruct
	Data *T   `json:"data,omitempty"`
	Type Type `json:"type,omitempty"`
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (t *metaInstruct) Bytes() []byte {
	t.Version = version.VersionTwo
	return instructToBytes(t)
}

func (t *metaInstruct) SetVersion(version string) *metaInstruct {
	t.Version = version
	return t
}

func (t *metaInstruct) JSON() ([]byte, error) {
	t.Version = version.VersionOne
	return json.Marshal(t)
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
		i.meta.Data = utils.Must(json.Marshal(i.Data))
		i.meta.Length = len(i.meta.Data)
	}
	return i.meta
}

func instructToBytes(c *metaInstruct) []byte {
	builder := flatbuffers.NewBuilder(0)
	_data := builder.CreateByteString(c.Data)
	_version := builder.CreateString(c.Version)
	instruct.InstructStart(builder)
	instruct.InstructAddData(builder, _data)
	instruct.InstructAddVersion(builder, _version)
	instruct.InstructAddType(builder, c.Type)
	builder.Finish(instruct.InstructEnd(builder))
	return builder.FinishedBytes()
}

// NewInstruct ...
// @param instruct.Type
// @return *Instruct[T]
func NewInstruct[T any]() *Instruct[T] {
	inst := new(Instruct[T])
	switch any(*new(T)).(type) {
	case ResourceData:
		inst.Type = instruct.TypeResource
	case CorrectData:
		inst.Type = instruct.TypeCorrect
	case ReportData:
		inst.Type = instruct.TypeReport
	}
	return inst
}

// getInstructType ...
// @param instruct.Type
// @return *Instruct[T]
func getInstructType(p Type) any {
	switch p {
	case instruct.TypeResource:
		return ResourceData{}
	case instruct.TypeCorrect:
		return CorrectData{}
	case instruct.TypeReport:
		return ReportData{}
	default:
		panic(ErrWrongInstructType)
	}
}
