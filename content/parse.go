package content

import (
	"encoding/json"
	"io"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/version"
)

// ParseJSONContent ...
// @Description: parse version 1 json data
// @param []byte
// @return *Content
// @return error
func ParseJSONContent(bytes []byte) (*Content, error) {
	logger.Println("print content json: ", string(bytes[:256]))
	var meta metaContent
	err := json.Unmarshal(bytes, &meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}

	ctnt := meta.content()
	ctnt.Message, err = parseRawMessage(meta.Message)
	logger.Println("print content parsed: ", string(ctnt.JSON()[:256]))
	return ctnt, err
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
		return nil, ErrWrongVersionType
	}

	ctnt := meta.content()
	ctnt.Message, err = parseRawMessage(meta.Message)
	logger.Println("print content parsed: ", string(ctnt.JSON()[:256]))
	return ctnt, err
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
		default:
			return nil, ErrWrongMessageType
		}
		msg.Version = CurrentDataVersion
		return msg, nil
	}
	return nil, nil
}

// ParserJSONContentWithV3
// @param []byte
// @return *Content
// @return error
func ParserJSONContentWithV3[T any](bytes []byte) (*ContentV3, error) {
	var meta metaContent
	err := json.Unmarshal(bytes, &meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}

	ctnt := meta.contentV3()
	//todo: parse message
	//ctnt.Message, err = parseRawMessageV3(meta.Message)
	return ctnt, err
}

// ParseContent ...
// @Description: parse version 2 flatbuffer data
// @param []byte
// @return retC
// @return errors
func ParseContent(bytes []byte) (retC *Content, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = Errorf("parse content error: %v", rerr)
		}
	}()
	c := content.GetRootAsContent(bytes, 0)
	if string(c.Version()) != version.VersionTwo {
		return nil, ErrWrongVersionType
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
		logger.Println("message is not empty", "last", message.Last(), "index", message.Index(), "size", len(message.Data()))
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
