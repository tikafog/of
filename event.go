package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(res *EventResult)

type Event interface {
	Trigger(n Name, r *EventRequest) error
	Callback(n Name, r *EventRequest) (<-chan EventResult, error)
	Register(from Name, event string, fn TypeEventFunc) error
}

type EventRequest struct {
	Receiver Name
	Type     string
	Args     []Arg
	Value    any
	Callback TypeEventCallbackFunc
}

type EventResult struct {
	Result any
	Error  error
}
