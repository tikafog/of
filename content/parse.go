package content

import (
	"encoding/json"
	"io"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/merr"
	"github.com/tikafog/of/version"
)

// ParseJSONContent ...
// @Description: parse version 1 json data
// @param []byte
// @return *Content
// @return error
func ParseJSONContent(bytes []byte) (*Content, error) {
	var meta metaContent
	err := json.Unmarshal(bytes, &meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, merr.Error(ErrWrongVersionType)
	}

	var content Content
	content.Message, err = parseRawMessage(meta.Message)
	return &content, err
}

// ParseJSONContentFromReader
// @param io.Reader
// @return *Content
// @return error
func ParseJSONContentFromReader(reader io.Reader) (*Content, error) {
	var meta metaContent
	err := json.NewDecoder(reader).Decode(&meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, merr.Error(ErrWrongVersionType)
	}
	var content Content
	content.Message, err = parseRawMessage(meta.Message)
	return &content, err
}

func parseRawMessage(raw json.RawMessage) (*Message, error) {
	var msg *Message
	if len(raw) != 0 {
		var mv metaMessageVersion
		if err := json.Unmarshal(raw, &mv); err != nil {
			return nil, err
		}
		switch mv.Version {
		case MessageV1Version:
			var v1 MessageV1
			if err := json.Unmarshal(raw, &v1); err != nil {
				return nil, err
			}
			msg = v1.current()
		case MessageV2Version:
			var v2 MessageV2
			if err := json.Unmarshal(raw, &v2); err != nil {
				return nil, err
			}
			msg = v2.current()
		}
		msg.Version = CurrentDataVersion
		return msg, nil
	}
	return nil, nil
}

//ParseContent ...
//@Description: parse version 2 flatbuffer data
//@param []byte
//@return retC
//@return errors
func ParseContent(bytes []byte) (retC *Content, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			//if e, ok := rerr.(error); ok && errors.IndexIs(ErrWrongVersionType, e) {
			//	errors = errors.Error(ErrWrongVersionType)
			//	return
			//}
			err = Errorf("parse content error: %v", rerr)
		}
	}()
	c := content.GetRootAsContent(bytes, 0)
	if string(c.Version()) != version.VersionTwo {
		return nil, merr.Error(ErrWrongVersionType)
	}
	return parseContent(c), err
}

func parseContent(cc *content.Content) *Content {
	//var node oc.Node
	message := cc.Message(nil)
	ext := make([]content.Ext, cc.ExtLength())
	var exts []Ext
	for i := 0; i < cc.ExtLength(); i++ {
		if cc.Ext(&ext[i], i) {
			exts = append(exts, Ext{
				ExtType: ext[i].Type(),
				Length:  len(ext[i].Data()),
				Data:    ext[i].Data(),
			})
		}
	}

	oc := NewContent(cc.Type())
	if message != nil {
		oc.Message = &Message{
			Last:    message.Last(),
			Index:   message.Index(),
			Version: int(message.Version()),
			Length:  len(message.Data()),
			Data:    message.Data(),
		}
	}
	oc.Exts = exts
	return oc
}
