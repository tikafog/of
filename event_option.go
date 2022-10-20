package of

import (
	"encoding/json"
)

type EventRequestOption struct {
	value    any
	args     []Arg
	data     []byte
	callback EventCallbackFunc
}

func (e EventRequestOption) Value() any {
	return e.value
}

func (e EventRequestOption) Args() []Arg {
	return e.args
}

func (e EventRequestOption) DecodeFromData(data any) error {
	return json.Unmarshal(e.data, data)
}

func (e EventRequestOption) Callback() EventCallbackFunc {
	return e.callback
}

type EventRequestOptions func(*EventRequestOption)

func EROValue(v any) EventRequestOptions {
	return func(e *EventRequestOption) {
		e.value = v
	}
}

func EROCallback(c EventCallbackFunc) EventRequestOptions {
	return func(e *EventRequestOption) {
		e.callback = c
	}
}

func EROArg(arg Arg) EventRequestOptions {
	return func(e *EventRequestOption) {
		e.args = append(e.args, arg)
	}
}

func EROArgs(args []Arg) EventRequestOptions {
	return func(e *EventRequestOption) {
		e.args = args
	}
}

func EROData(data any) EventRequestOptions {
	return func(e *EventRequestOption) {
		e.data, _ = json.Marshal(data)
	}
}
