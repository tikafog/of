package instruct

import (
	"encoding/json"
	"io"

	"github.com/tikafog/of/buffers/instruct"
	"github.com/tikafog/of/merr"
	"github.com/tikafog/of/version"
)

// ParseJSONInstruct ...
// @Description: parse version 1 json data
// @param []byte
// @return *metaInstruct
// @return error
func ParseJSONInstruct(bytes []byte) (any, error) {
	var meta metaInstruct
	err := json.Unmarshal(bytes, &meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}
	return parseMetaInstruct(&meta)
}

func ParseJSONInstructFromReader(reader io.Reader) (any, error) {
	var meta metaInstruct
	err := json.NewDecoder(reader).Decode(&meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, merr.Error(ErrWrongVersionType)
	}
	return parseMetaInstruct(&meta)
}

//ParseInstruct ...
//@Description: parse version 2 flatbuffer data
//@param []byte
//@return retC
//@return errors
func ParseInstruct(bytes []byte) (retC any, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = Errorf("parse instruct error: %v", rerr)
		}
	}()
	c := instruct.GetRootAsInstruct(bytes, 0)
	if string(c.Version()) != version.VersionTwo {
		return nil, merr.Error(ErrWrongVersionType)
	}
	meta := instructToMetaInstruct(c)
	return parseMetaInstruct(meta)
}

func parseInstructT[T any](meta *metaInstruct) (*Instruct[T], error) {
	var inst Instruct[T]
	inst.Type = meta.Type
	//inst.Version = meta.Version
	if meta.Length != 0 && inst.Data == nil {
		inst.Data = new(T)
		err := json.Unmarshal(meta.Data, inst.Data)
		return &inst, err
	}

	return &inst, nil
}

type instructBuffer interface {
	decodeInstruct(inst *instruct.Instruct) error
	encodeInstruct() []byte
}

func instructToMetaInstruct(cc *instruct.Instruct) *metaInstruct {
	var inst metaInstruct
	inst.Version = string(cc.Version())
	inst.Type = cc.Type()
	inst.Data = cc.Data()
	inst.Length = len(inst.Data)
	return &inst
}

func parseMetaInstruct(meta *metaInstruct) (any, error) {
	var inst metaParser
	switch meta.Type {
	case instruct.TypeResource:
		inst = new(Instruct[ResourceData])
	default:
		return nil, ErrWrongInstructType
	}
	err := inst.parseMetaInstruct(meta)
	return inst, err
}
