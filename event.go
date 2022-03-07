package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(v interface{}, err error)

type Event interface {
	Event(n Name, r *EventRequest) error
	EventWithCallback(n Name, r *EventRequest) (<-chan EventResult, error)
}

type EventRequest struct {
	Receiver Name
	Type     string
	Value    interface{}
	Callback TypeEventCallbackFunc
}

type EventResult struct {
	Result interface{}
	Error  error
}
