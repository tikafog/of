// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package of

import (
	"fmt"

	"github.com/tikafog/errors"
)

const (
	// ProtocolAnswer is a Protocol of type Answer.
	ProtocolAnswer Protocol = iota
	// ProtocolGossip is a Protocol of type Gossip.
	ProtocolGossip
	// ProtocolDataSwap is a Protocol of type DataSwap.
	ProtocolDataSwap
	// ProtocolMessage is a Protocol of type Message.
	ProtocolMessage
	// ProtocolMessageV2 is a Protocol of type MessageV2.
	ProtocolMessageV2
	// ProtocolInstruct is a Protocol of type Instruct.
	ProtocolInstruct
	// ProtocolAdapter is a Protocol of type Adapter.
	ProtocolAdapter
	// ProtocolMax is a Protocol of type Max.
	ProtocolMax
)

var ErrInvalidProtocol = errors.New("not a valid Protocol")

const _ProtocolName = "AnswerGossipDataSwapMessageMessageV2InstructAdapterMax"

var _ProtocolMap = map[Protocol]string{
	ProtocolAnswer:    _ProtocolName[0:6],
	ProtocolGossip:    _ProtocolName[6:12],
	ProtocolDataSwap:  _ProtocolName[12:20],
	ProtocolMessage:   _ProtocolName[20:27],
	ProtocolMessageV2: _ProtocolName[27:36],
	ProtocolInstruct:  _ProtocolName[36:44],
	ProtocolAdapter:   _ProtocolName[44:51],
	ProtocolMax:       _ProtocolName[51:54],
}

// String implements the Stringer interface.
func (x Protocol) String() string {
	if str, ok := _ProtocolMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Protocol(%d)", x)
}

var _ProtocolValue = map[string]Protocol{
	_ProtocolName[0:6]:   ProtocolAnswer,
	_ProtocolName[6:12]:  ProtocolGossip,
	_ProtocolName[12:20]: ProtocolDataSwap,
	_ProtocolName[20:27]: ProtocolMessage,
	_ProtocolName[27:36]: ProtocolMessageV2,
	_ProtocolName[36:44]: ProtocolInstruct,
	_ProtocolName[44:51]: ProtocolAdapter,
	_ProtocolName[51:54]: ProtocolMax,
}

// ParseProtocol attempts to convert a string to a Protocol.
func ParseProtocol(name string) (Protocol, error) {
	if x, ok := _ProtocolValue[name]; ok {
		return x, nil
	}
	return Protocol(0), fmt.Errorf("%s is %w", name, ErrInvalidProtocol)
}
