package utils

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/tikafog/of/buffers/message"
)

func NewBufferMessage(id string, topic string, last int64, data []byte) []byte {
	builder := flatbuffers.NewBuilder(1024)
	_id := builder.CreateString(id)
	_topic := builder.CreateString(topic)
	_data := builder.CreateByteString(data)

	message.MessageStart(builder)
	message.MessageAddId(builder, _id)
	message.MessageAddTopic(builder, _topic)
	message.MessageAddLast(builder, last)
	message.MessageAddData(builder, _data)
	builder.Finish(message.MessageEnd(builder))

	return builder.FinishedBytes()
}

func ParseBufferMessage(data []byte) (msg *message.Message) {
	defer func() {
		if err := recover(); err != nil {
			msg = nil
		}
	}()
	return message.GetRootAsMessage(data, 0)
}
