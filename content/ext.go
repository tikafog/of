package content

import (
	"github.com/tikafog/of/buffers/content"
)

// ExtConverter ...
// @Description:
type ExtConverter interface {
	ExtType() content.ExtType
	MarshalData() ([]byte, error)
	UnmarshalData(data []byte) error
}

// ExtType ...
type ExtType = content.ExtType

// Ext ...
// @Description:
type Ext struct {
	ExtType content.ExtType `json:"ext_type,omitempty"`
	Length  int             `json:"length,omitempty"`
	Data    []byte          `json:"data,omitempty"`
}

// ParseExt ...
// @Description:
// @param Ext
// @param interface{}
// @return error
func ParseExt(ext Ext, v ExtConverter) error {
	if ext.ExtType != v.ExtType() {
		return merr.Errorf("extension type %v is different with %v", v.ExtType(), ext.ExtType)
	}

	if ext.Length == 0 {
		return nil
	}
	return v.UnmarshalData(ext.Data)
}

// ParseExtConverter ...
// @Description:
// @param Ext
// @param interface{}
// @return error
func ParseExtConverter(ext Ext, v interface{}) error {
	vv, b := v.(ExtConverter)
	if !b {
		return merr.Error(ErrWrongExtType)
	}
	if ext.ExtType != vv.ExtType() {
		return merr.Errorf("extension type %v is different with %v", vv.ExtType(), ext.ExtType)
	}

	if ext.Length == 0 {
		return nil
	}
	return vv.UnmarshalData(ext.Data)
}

// MakeExtConverter ...
// @Description:
// @param interface{}
// @return Ext
// @return error
func MakeExtConverter(v interface{}) (Ext, error) {
	var ext Ext
	var err error
	vv, b := v.(ExtConverter)
	if !b {
		return ext, merr.Error(ErrWrongExtType)
	}
	ext.ExtType = vv.ExtType()
	ext.Data, err = vv.MarshalData()
	if err != nil {
		return ext, merr.WrapString(err, "marshall data failed")
	}
	ext.Length = len(ext.Data)
	return ext, nil
}

// MakeExt ...
// @Description:
// @param interface{}
// @return Ext
// @return error
func MakeExt(v ExtConverter) (Ext, error) {
	var ext Ext
	var err error
	ext.ExtType = v.ExtType()
	ext.Data, err = v.MarshalData()
	if err != nil {
		return ext, merr.WrapString(err, "marshall data failed")
	}
	ext.Length = len(ext.Data)
	return ext, nil
}
