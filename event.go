package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(res *EventResult)

type Event interface {
	Trigger(self Name, r *EventRequest) error
	TriggerAsync(self Name, r *EventRequest) error
	//Callback(self Name, r *EventRequest) (<-chan EventResult, error)
	Register(from Name, event string, fn TypeEventFunc) error
}

type EventRequest struct {
	Receiver Name
	Type     string
	Args     []Arg
	Value    any
	//Callback TypeEventCallbackFunc
}

type EventResult struct {
	Result any
	Error  error
}
