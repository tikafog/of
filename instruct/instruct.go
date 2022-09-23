package instruct

import (
	"encoding/json"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tikafog/of/buffers/instruct"
	"github.com/tikafog/of/utils"
	"github.com/tikafog/of/version"
)

const CurrentDataVersion = 2

const (
// ErrUnsupportedExtType ...
//ErrUnsupportedExtType = errors.New("unsupported ext type")
// ErrMustBePointer ...
//ErrMustBePointer = errors.New("the interface parameter must be a pointer type")
//WrongVersionType = "wrong version type"
)

//var (
//	ErrWrongVersionType = Error(WrongVersionType)
//)

// Type ...
type Type = instruct.Type

//type Instruct[T] struct {
//	Version string          `json:"version,omitempty"`
//	Data    json.RawMessage `json:"data,omitempty"`
//	Type    Type            `json:"type,omitempty"`
//}

type metaParser interface {
	parseMeta(m *metaInstruct) error
}

type metaInstruct struct {
	Version string          `json:"version,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Type    Type            `json:"type,omitempty"`
}

// Instruct ...
type Instruct[T any] struct {
	meta    *metaInstruct
	Version string `json:"version,omitempty"`
	Data    *T     `json:"data,omitempty"`
	Type    Type   `json:"type,omitempty"`
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (t *metaInstruct) Bytes() []byte {
	t.Version = version.VersionTwo
	return instructToBytes(t)
}

func (t *Instruct[T]) decodeInstruct(inst *instruct.Instruct) error {
	if inst == nil || len(inst.Data()) == 0 {
		return nil
	}
	t.Version = string(inst.Version())
	t.Type = inst.Type()
	t.Data = new(T)
	return json.Unmarshal(inst.Data(), &t.Data)
}

func (t *Instruct[T]) encodeInstruct() []byte {
	return instructToBytes(&metaInstruct{
		Version: version.VersionTwo,
		Data:    utils.Must(t.MarshalData()),
		Type:    t.Type,
	})
}

// MarshalData ...
// @Description:
// @receiver ContentReport
// @return data
func (t *Instruct[T]) MarshalData() ([]byte, error) {
	return json.Marshal(t)
}

// UnmarshalData ...
// @Description:
// @receiver *Instruct
// @param []byte
// @return error
func (t *Instruct[T]) UnmarshalData(data []byte) error {
	return json.Unmarshal(data, t)
}

// SetData
// @receiver *Instruct[T]
// @param []byte
// @return *Instruct[T]
func (t *Instruct[T]) SetData(data *T) *Instruct[T] {
	t.Data = data
	return t
}

// SetType
// @receiver *Instruct[T]
// @param instruct.Type
// @return *Instruct[T]
func (t *Instruct[T]) SetType(p instruct.Type) *Instruct[T] {
	t.Type = p
	return t
}

// JSON ...
// @Description:
// @receiver Instruct[T]
// @return []byte
// @return error
func (t *Instruct[T]) JSON() ([]byte, error) {
	t.Version = version.VersionOne
	return json.Marshal(t)
}

// MustJSON ...
// @Description:
// @receiver Instruct[T]
// @return []byte
// @return error
func (t *Instruct[T]) MustJSON() []byte {
	data, _ := t.JSON()
	//if errors != nil {
	//	panic(errors)
	//}
	return data
}

// FinishBytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
// Decrypted use Bytes instead
func (t *Instruct[T]) FinishBytes() []byte {
	return t.Bytes()
}

// Bytes ...
// @Description:
// @receiver Instruct[T]
// @return []byte
func (t *Instruct[T]) Bytes() []byte {
	t.Version = version.VersionTwo
	return instructToBytes(&metaInstruct{
		Version: version.VersionTwo,
		Type:    t.Type,
		Data:    utils.Must(json.Marshal(t.Data)),
	})
}

func (t *Instruct[T]) parseMeta(m *metaInstruct) error {
	t.meta = m
	t.Type = m.Type
	t.Version = m.Version
	if len(m.Data) != 0 && t.Data == nil {
		t.Data = new(T)
		return json.Unmarshal(m.Data, t.Data)
	}
	return nil
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
	case Resource:
		inst.Type = instruct.TypeResource
	}
	return inst
}

// NewInstruct ...
// @param instruct.Type
// @return *Instruct[T]
func getInstructType(p Type) any {
	switch p {
	case instruct.TypeResource:
		return Resource{}
	default:
		panic(ErrWrongInstructType)
	}
}
