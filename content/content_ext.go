package content

import (
	"errors"
	"fmt"

	"github.com/tikafog/of/buffers/content"
)

// ExtConvertable ...
// @Description:
type ExtConvertable interface {
	ExtType() content.ExtType
	MarshalData() ([]byte, error)
	UnmarshalData(data []byte) error
}

// ErrWrongExtType ...
var ErrWrongExtType = errors.New("wrong extension type")

// ParseExt ...
// @Description:
// @param Ext
// @param interface{}
// @return error
func ParseExt(ext Ext, v interface{}) error {
	vv, b := v.(ExtConvertable)
	if !b {
		return ErrWrongExtType
	}
	if ext.ExtType != vv.ExtType() {
		return fmt.Errorf("given extension type %v is different with %v", vv.ExtType(), ext.ExtType)
	}

	if ext.Length == 0 {
		return nil
	}
	return vv.UnmarshalData(ext.Data)
}

// MakeExt ...
// @Description:
// @param interface{}
// @return Ext
// @return error
func MakeExt(v interface{}) (Ext, error) {
	var ext Ext
	var err error
	vv, b := v.(ExtConvertable)
	if !b {
		return ext, ErrWrongExtType
	}
	ext.ExtType = vv.ExtType()
	ext.Data, err = vv.MarshalData()
	if err != nil {
		return ext, err
	}
	ext.Length = len(ext.Data)
	return ext, nil
}
