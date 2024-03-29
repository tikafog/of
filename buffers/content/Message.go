// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package content

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Message struct {
	_tab flatbuffers.Table
}

func GetRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Message{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMessage(buf []byte, offset flatbuffers.UOffsetT) *Message {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Message{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Message) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Message) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Message) Index() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Message) MutateIndex(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Message) Last() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Message) MutateLast(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *Message) Version() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Message) MutateVersion(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Message) Data() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func MessageAddIndex(builder *flatbuffers.Builder, index int64) {
	builder.PrependInt64Slot(0, index, 0)
}
func MessageAddLast(builder *flatbuffers.Builder, last int64) {
	builder.PrependInt64Slot(1, last, 0)
}
func MessageAddVersion(builder *flatbuffers.Builder, version int32) {
	builder.PrependInt32Slot(2, version, 0)
}
func MessageAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(data), 0)
}
func MessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
