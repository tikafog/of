package of

import (
	"encoding/json"
)

type EventRequestOption interface {
	Value() any
	Args() []Arg
	DecodeFromData(data any) error
	Callback(res *EventResult)
}

type eventRequestOption struct {
	value    any
	args     []Arg
	data     []byte
	callback EventCallbackFunc
}

func (e eventRequestOption) Value() any {
	return e.value
}

func (e eventRequestOption) Args() []Arg {
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

func EROArg(arg Arg) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.args = append(e.args, arg)
	}
}

func EROArgs(args []Arg) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.args = args
	}
}

func EROData(data any) EventRequestOptions {
	return func(e *eventRequestOption) {
		e.data, _ = json.Marshal(data)
	}
}
