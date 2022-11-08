package instruct

import (
	"encoding/json"
	"io"

	"github.com/tikafog/of/buffers/instruct"
	"github.com/tikafog/of/version"
)

// ParseJSONInstruct ...
// @Description: parse version 1 json data
// @param []byte
// @return *metaInstruct
// @return error
func ParseJSONInstruct(bytes []byte) (Instructor, error) {
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

// ParseJSONInstructFromReader ...
// @param io.Reader
// @return Instructor
// @return error
func ParseJSONInstructFromReader(reader io.Reader) (Instructor, error) {
	var meta metaInstruct
	err := json.NewDecoder(reader).Decode(&meta)
	if err != nil {
		return nil, err
	}
	if string(meta.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}
	return parseMetaInstruct(&meta)
}

// ParseInstruct ...
// @Description: parse version 2 flatbuffer data
// @param []byte
// @return retC
// @return errors
func ParseInstruct(bytes []byte) (retC Instructor, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = Errorf("parse instruct error: %v", rerr)
		}
	}()
	c := instruct.GetRootAsInstruct(bytes, 0)
	if string(c.Version()) != version.VersionTwo {
		return nil, ErrWrongVersionType
	}
	meta := instructToMetaInstruct(c)
	return parseMetaInstruct(meta)
}

func parseInstructT[T DataSource](meta *metaInstruct) (*Instruct[T], error) {
	var inst Instruct[T]
	if meta.Length != 0 && meta.Data == nil {
		inst.dataSource = new(T)
		err := json.Unmarshal(meta.Data, inst.dataSource)
		return &inst, err
	}
	return &inst, nil
}

func instructToMetaInstruct(cc *instruct.Instruct) *metaInstruct {
	var inst metaInstruct
	inst.Version = string(cc.Version())
	inst.Type = cc.Type()
	inst.Data = cc.Data()
	inst.Last = cc.Last()
	inst.Length = len(inst.Data)
	return &inst
}

func parseMetaInstruct(meta *metaInstruct) (Instructor, error) {
	var inst metaParser
	switch meta.Type {
	case instruct.TypeResource:
		inst = NewInstruct[ResourceData]()
	case instruct.TypeCorrect:
		inst = NewInstruct[CorrectData]()
	case instruct.TypeReport:
		inst = NewInstruct[ReportData]()
	default:
		return nil, ErrWrongInstructType
	}
	err := inst.parseMetaInstruct(meta)
	return inst.(Instructor), err
}
