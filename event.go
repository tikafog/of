package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(res *EventResult)

type Event interface {
	Event(n Name, r *EventRequest) error
	EventWithCallback(n Name, r *EventRequest) (<-chan EventResult, error)
	RegisterEventHandler(from Name, fn TypeEventFunc) error
}

type EventRequest struct {
	Receiver Name
	Type     string
	Args     []Arg
	Value    interface{}
	Callback TypeEventCallbackFunc
}

type EventResult struct {
	Result interface{}
	Error  error
}
