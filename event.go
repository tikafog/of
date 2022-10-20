package of

type TypeEventFunc = func(r *EventRequest) error
type EventCallbackFunc = func(res *EventResult)

type EventTrigger interface {
	Trigger(key EventKey, opts ...EventRequestOptions) error
}

type Event interface {
	EventTrigger
	Register(key EventKey, fn TypeEventFunc) error
}

type EventRegister interface {
	RegisterEvent(Event) error
}

type EventRequest struct {
	EventKey
	*EventRequestOption
}

type EventResult struct {
	Result any
	Error  error
}
