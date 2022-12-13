package of

import (
	"encoding/json"
)

type EventRequestOption interface {
	Value() any
	Args() Args
	DecodeFromData(data any) error
	Callback(res *EventResult)
}

type eventRequestOption struct {
	receiver Name
	value    any
	args     Args
	data     []byte
	callback EventCallbackFunc
}

func (e *eventRequestOption) Receiver() Name {
	return e.receiver
}

func (e *eventRequestOption) SetReceiver(receiver Name) {
	e.receiver = receiver
}

func (e eventRequestOption) Value() any {
	return e.value
}

func (e eventRequestOption) Args() Args {
	return e.args
}

func (e eventRequestOption) DecodeFromData(data any) error {
	return json.Unmarshal(e.data, data)
}

func (e eventRequestOption) Callback(res *EventResult) {
	if e.callback != nil {
		e.callback(res)
	}
}

type EventRequestOptions func(*eventRequestOption)

func EROValue(v any) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.value = v
	}
}

func EROCallback(c EventCallbackFunc) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.callback = c
	}
}

func EROArg(arg KeyArg) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.args.Add(arg)
	}
}

func EROArgs(args Args) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.args = args
	}
}

func EROData(data any) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.data, _ = json.Marshal(data)
	}
}

func EROReceiver(name Name) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.receiver = name
	}
}

func ParseERO(opts ...EventRequestOptions) EventRequestOption {
	op := initOpts()
	for i := range opts {
		opts[i](op)
	}
	return op
}

func initOpts() *eventRequestOption {
	return &eventRequestOption{
		args: NewArgs(),
	}
}
