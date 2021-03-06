// Code generated by go-enum DO NOT EDIT.
// Version: 0.3.10
// Revision: 07a5b318d9ef317345f0cfca0c9347eda0e2bfc4
// Build Date: 2022-01-23T20:31:33Z
// Built By: goreleaser

package content

import (
	"fmt"
)

const (
	// NodeStateInitialized is a NodeState of type Initialized.
	NodeStateInitialized NodeState = iota
	// NodeStateStart is a NodeState of type Start.
	NodeStateStart
	// NodeStateOnline is a NodeState of type Online.
	NodeStateOnline
	// NodeStateOffline is a NodeState of type Offline.
	NodeStateOffline
	// NodeStateMax is a NodeState of type Max.
	NodeStateMax
)

const _NodeStateName = "initializedstartonlineofflinemax"

var _NodeStateMap = map[NodeState]string{
	NodeStateInitialized: _NodeStateName[0:11],
	NodeStateStart:       _NodeStateName[11:16],
	NodeStateOnline:      _NodeStateName[16:22],
	NodeStateOffline:     _NodeStateName[22:29],
	NodeStateMax:         _NodeStateName[29:32],
}

// String implements the Stringer interface.
func (x NodeState) String() string {
	if str, ok := _NodeStateMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NodeState(%d)", x)
}

var _NodeStateValue = map[string]NodeState{
	_NodeStateName[0:11]:  NodeStateInitialized,
	_NodeStateName[11:16]: NodeStateStart,
	_NodeStateName[16:22]: NodeStateOnline,
	_NodeStateName[22:29]: NodeStateOffline,
	_NodeStateName[29:32]: NodeStateMax,
}

// ParseNodeState attempts to convert a string to a NodeState.
func ParseNodeState(name string) (NodeState, error) {
	if x, ok := _NodeStateValue[name]; ok {
		return x, nil
	}
	return NodeState(0), fmt.Errorf("%s is not a valid NodeState", name)
}

const (
	// NodeTypeServer is a NodeType of type Server.
	NodeTypeServer NodeType = iota
	// NodeTypeAdapter is a NodeType of type Adapter.
	NodeTypeAdapter
	// NodeTypeBox is a NodeType of type Box.
	NodeTypeBox
	// NodeTypeMobile is a NodeType of type Mobile.
	NodeTypeMobile
	// NodeTypeMax is a NodeType of type Max.
	NodeTypeMax
)

const _NodeTypeName = "serveradapterboxmobilemax"

var _NodeTypeMap = map[NodeType]string{
	NodeTypeServer:  _NodeTypeName[0:6],
	NodeTypeAdapter: _NodeTypeName[6:13],
	NodeTypeBox:     _NodeTypeName[13:16],
	NodeTypeMobile:  _NodeTypeName[16:22],
	NodeTypeMax:     _NodeTypeName[22:25],
}

// String implements the Stringer interface.
func (x NodeType) String() string {
	if str, ok := _NodeTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NodeType(%d)", x)
}

var _NodeTypeValue = map[string]NodeType{
	_NodeTypeName[0:6]:   NodeTypeServer,
	_NodeTypeName[6:13]:  NodeTypeAdapter,
	_NodeTypeName[13:16]: NodeTypeBox,
	_NodeTypeName[16:22]: NodeTypeMobile,
	_NodeTypeName[22:25]: NodeTypeMax,
}

// ParseNodeType attempts to convert a string to a NodeType.
func ParseNodeType(name string) (NodeType, error) {
	if x, ok := _NodeTypeValue[name]; ok {
		return x, nil
	}
	return NodeType(0), fmt.Errorf("%s is not a valid NodeType", name)
}
