package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(res *EventResult)

type Event interface {
	Trigger(self Name, r *EventRequest) error
	TriggerAsync(self Name, r *EventRequest) error
	Register(from Name, event string, fn TypeEventFunc) error
	//EventKey(name string) uint64
	//Register(from Name, event uint64, fn TypeEventFunc) error
}

type EventRegister interface {
	Register(Event) error
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
