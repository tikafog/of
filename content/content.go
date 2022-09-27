package content

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/utils"
	"github.com/tikafog/of/version"
)

const (
// ErrUnsupportedExtType ...
// ErrUnsupportedExtType = errors.New("unsupported ext type")
// ErrMustBePointer ...
// ErrMustBePointer = errors.New("the interface parameter must be a pointer type")
// WrongVersionType = "wrong version type"
)

//var (
//	ErrWrongVersionType = Error(WrongVersionType)
//)

// Type ...
type Type = content.Type

type metaContent struct {
	Version string          `json:"version,omitempty"`
	From    string          `json:"from,omitempty"`
	Message json.RawMessage `json:"message,omitempty"`
	Exts    []Ext           `json:"ext,omitempty"`
	Type    content.Type    `json:"type,omitempty"`
}

// Content Content
type Content struct {
	meta    *metaContent
	From    string
	Message *Message
	Exts    []Ext
	Type    content.Type
}

func (m metaContent) content() *Content {
	return &Content{
		meta: &m,
		From: m.From,
		Exts: m.Exts,
		Type: m.Type,
	}
}

// SetExts
// @receiver *Content
// @param ...Ext
// @return *Content
func (c *Content) SetExts(exts ...Ext) *Content {
	c.Exts = exts
	return c
}

// AddExts
// @receiver *Content
// @param ...Ext
// @return *Content
func (c *Content) AddExts(exts ...Ext) *Content {
	c.Exts = append(c.Exts, exts...)
	return c
}

// SetMessage
// @receiver *Content
// @param *Message
// @return *Content
func (c *Content) SetMessage(m *Message) *Content {
	c.Message = m
	return c
}

// NewMessage
// @receiver *Content
// @param []byte
// @return *Content
func (c *Content) NewMessage(data []byte) *Content {
	if data != nil {
		c.Message = NewContentMessage(data)
	}
	return c
}

// NewMessageDetail
// @receiver *Content
// @param []byte
// @param int64
// @param int64
// @return *Content
func (c *Content) NewMessageDetail(data []byte, index, last int64) *Content {
	if data != nil {
		c.Message = NewContentMessageWithDetail(data, index, last)
	}
	return c
}

// NewMessageLast
// @receiver *Content
// @param int64
// @return *Content
func (c *Content) NewMessageLast(last int64) *Content {
	c.Message = NewContentMessageLast(last)
	return c
}

// NewMessageAndLast
// @receiver *Content
// @param []byte
// @param int64
// @return *Content
func (c *Content) NewMessageAndLast(data []byte, last int64) *Content {
	c.Message = NewContentMessageLast(last)
	return c
}

// SetType
// @receiver *Content
// @param content.Type
// @return *Content
func (c *Content) SetType(p content.Type) *Content {
	c.Type = p
	return c
}

// SetFrom
// @receiver *Content
// @param string
// @return *Content
func (c *Content) SetFrom(s string) *Content {
	c.From = s
	return c
}

// JSON ...
// @Description:
// @receiver Content
// @return []byte
// @return error
func (c *Content) JSON() []byte {
	if c.meta == nil {
		var message []byte
		if c.Message != nil {
			message = utils.Must(json.Marshal(c.Message.v1()))
		}
		c.meta = &metaContent{
			Version: "",
			From:    c.From,
			Message: message,
			Exts:    c.Exts,
			Type:    c.Type,
		}
	}
	c.meta.Version = version.VersionOne
	return utils.Must(json.Marshal(c.meta))
}

// JSONV2 ...
// @Description:
// @receiver Content
// @return []byte
// @return error
func (c *Content) JSONV2() []byte {
	if c.meta == nil {
		var message []byte
		if c.Message != nil {
			message = utils.Must(json.Marshal(c.Message))
		}
		c.meta = &metaContent{
			Version: "",
			From:    c.From,
			Message: message,
			Exts:    c.Exts,
			Type:    c.Type,
		}
	}
	c.meta.Version = version.VersionOne
	return utils.Must(json.Marshal(c.meta))
}

// Clear
// @receiver *Content
func (c *Content) Clear() {
	c.Message = EmptyMessage
	c.Exts = []Ext{}
}

// FinishBytes ...
// @Description:
// @receiver Content
// @return []byte
// Decrypted use Bytes instead
func (c *Content) FinishBytes() []byte {
	return c.Bytes()
}

// Bytes ...
// @Description:
// @receiver Content
// @return []byte
func (c *Content) Bytes() []byte {
	//if c.meta == nil {
	//	c.meta = &metaContent{
	//		Version: "",
	//		From:    c.From,
	//		Message: utils.Must(json.Marshal(c.Message)),
	//		Exts:    c.Exts,
	//		Type:    c.Type,
	//	}
	//}
	//c.meta.Version = version.VersionTwo
	return contentToBytes(c, c.Message.IsEmpty())
}

func contentToBytes(c *Content, has bool) []byte {
	builder := flatbuffers.NewBuilder(0)

	var _message flatbuffers.UOffsetT
	if !has {
		_dataM := builder.CreateByteString(c.Message.Data)
		content.MessageStart(builder)
		content.MessageAddLast(builder, c.Message.Last)
		content.MessageAddIndex(builder, c.Message.Index)
		content.MessageAddVersion(builder, int32(c.Message.Version))
		content.MessageAddData(builder, _dataM)
		_message = content.MessageEnd(builder)
	}
	var _exts []flatbuffers.UOffsetT
	for i := range c.Exts {
		_dataE := builder.CreateByteString(c.Exts[i].Data)
		content.ExtStart(builder)
		content.ExtAddType(builder, c.Exts[i].ExtType)
		content.ExtAddData(builder, _dataE)
		_exts = append(_exts, content.ExtEnd(builder))
	}
	content.ContentStartExtVector(builder, len(_exts))
	for i := range _exts {
		builder.PrependUOffsetT(_exts[i])
	}
	_extsVec := builder.EndVector(len(_exts))

	_version := builder.CreateString(version.VersionTwo)
	_from := builder.CreateString(c.From)
	content.ContentStart(builder)
	content.ContentAddExt(builder, _extsVec)
	if !has {
		content.ContentAddMessage(builder, _message)
	}
	content.ContentAddFrom(builder, _from)
	content.ContentAddVersion(builder, _version)
	content.ContentAddType(builder, c.Type)
	builder.Finish(content.ContentEnd(builder))
	return builder.FinishedBytes()
}

// NewContentWithMessage ...
// @param content.Type
// @param *Message
// @return *Content
func NewContentWithMessage(p content.Type, message *Message) *Content {
	return &Content{
		Type:    p,
		Message: message,
	}
}

// NewContentWithExts
// @param content.Type
// @param ...Ext
// @return *Content
func NewContentWithExts(p content.Type, exts ...Ext) *Content {
	return &Content{
		Type: p,
		Exts: exts,
	}
}

// NewContentWithType ...
// @param content.Type
// @return *Content
func NewContentWithType(p content.Type) *Content {
	return &Content{
		Type: p,
	}
}

// NewContent ...
// @param content.Type
// @return *Content
func NewContent(p content.Type) *Content {
	return &Content{
		Type: p,
	}
}

// NewTypeContent
// @param content.Type
// @param ...func(content *Content)
// @return *Content
// Decrypted: to be used NewContent
func NewTypeContent(tp content.Type, fns ...func(content *Content)) *Content {
	_content := Content{
		Type: tp,
	}
	for i := range fns {
		fns[i](&_content)
	}
	return &_content
}

// bytesToStringArray ...
// @Description: convert bytes to string array
// @param []byte
// @return []string
func bytesToStringArray(data []byte) []string {
	buf := bytes.NewBuffer(data)
	addr := struct {
		number int32
		lens   []int32
	}{}
	_ = binary.Read(buf, binary.BigEndian, &addr.number)
	addr.lens = make([]int32, addr.number)
	var ret []string
	for i := int32(0); i < addr.number; i++ {
		_ = binary.Read(buf, binary.BigEndian, &addr.lens[i])
		ret = append(ret, string(buf.Next(int(addr.lens[i]))))
	}
	return ret
}

// stringArrayToBytes ...
// @Description: convert string array to bytes
// @param ...string
// @return []byte
func stringArrayToBytes(strs ...string) []byte {
	buf := bytes.NewBuffer(nil)
	number := int32(len(strs))
	_ = binary.Write(buf, binary.BigEndian, number)
	var size int32
	for i := range strs {
		size = int32(len(strs[i]))
		_ = binary.Write(buf, binary.BigEndian, size)
		buf.WriteString(strs[i])
	}
	return buf.Bytes()
}
