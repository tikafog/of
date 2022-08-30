// Code generated by go-enum DO NOT EDIT.
// Version: 0.3.10
// Revision: 07a5b318d9ef317345f0cfca0c9347eda0e2bfc4
// Build Date: 2022-01-23T20:31:33Z
// Built By: goreleaser

package of

import (
	"fmt"
)

const (
	// ErrCoreIsNil is a Err of type Core Is Nil.
	ErrCoreIsNil Err = iota
	// ErrFatherNotFound is a Err of type Father Not Found.
	ErrFatherNotFound
	// ErrNoDataFound is a Err of type No Data Found.
	ErrNoDataFound
	// ErrChannelClosed is a Err of type Channel Closed.
	ErrChannelClosed
	// ErrSkipOldData is a Err of type Skip Old Data.
	ErrSkipOldData
)

const _ErrName = "core is nilfather not foundno data foundchannel closedskip old data"

var _ErrMap = map[Err]string{
	ErrCoreIsNil:      _ErrName[0:11],
	ErrFatherNotFound: _ErrName[11:27],
	ErrNoDataFound:    _ErrName[27:40],
	ErrChannelClosed:  _ErrName[40:54],
	ErrSkipOldData:    _ErrName[54:67],
}

// String implements the Stringer interface.
func (x Err) String() string {
	if str, ok := _ErrMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Err(%d)", x)
}

var _ErrValue = map[string]Err{
	_ErrName[0:11]:  ErrCoreIsNil,
	_ErrName[11:27]: ErrFatherNotFound,
	_ErrName[27:40]: ErrNoDataFound,
	_ErrName[40:54]: ErrChannelClosed,
	_ErrName[54:67]: ErrSkipOldData,
}

// ParseErr attempts to convert a string to a Err.
func ParseErr(name string) (Err, error) {
	if x, ok := _ErrValue[name]; ok {
		return x, nil
	}
	return Err(0), fmt.Errorf("%s is not a valid Err", name)
}
