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
	// ResourceTypeLocal is a ResourceType of type Local.
	ResourceTypeLocal ResourceType = iota
	// ResourceTypeRemote is a ResourceType of type Remote.
	ResourceTypeRemote
	// ResourceTypeHash is a ResourceType of type Hash.
	ResourceTypeHash
	// ResourceTypeRemoteHash is a ResourceType of type Remote Hash.
	ResourceTypeRemoteHash
	// ResourceTypeMax is a ResourceType of type Max.
	ResourceTypeMax
)

const _ResourceTypeName = "localremotehashremote hashmax"

var _ResourceTypeMap = map[ResourceType]string{
	ResourceTypeLocal:      _ResourceTypeName[0:5],
	ResourceTypeRemote:     _ResourceTypeName[5:11],
	ResourceTypeHash:       _ResourceTypeName[11:15],
	ResourceTypeRemoteHash: _ResourceTypeName[15:26],
	ResourceTypeMax:        _ResourceTypeName[26:29],
}

// String implements the Stringer interface.
func (x ResourceType) String() string {
	if str, ok := _ResourceTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ResourceType(%d)", x)
}

var _ResourceTypeValue = map[string]ResourceType{
	_ResourceTypeName[0:5]:   ResourceTypeLocal,
	_ResourceTypeName[5:11]:  ResourceTypeRemote,
	_ResourceTypeName[11:15]: ResourceTypeHash,
	_ResourceTypeName[15:26]: ResourceTypeRemoteHash,
	_ResourceTypeName[26:29]: ResourceTypeMax,
}

// ParseResourceType attempts to convert a string to a ResourceType.
func ParseResourceType(name string) (ResourceType, error) {
	if x, ok := _ResourceTypeValue[name]; ok {
		return x, nil
	}
	return ResourceType(0), fmt.Errorf("%s is not a valid ResourceType", name)
}