package of

type EventFunc = func(r *EventRequest) error
type EventCallbackFunc = func(res *EventResult)

type EventTrigger interface {
	Trigger(key EventKey, opts ...EventRequestOptions) error
}

type Event interface {
	EventTrigger
	Register(key EventKey, fn EventFunc) error
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

func parseERO(opts []EventRequestOptions) *EventRequestOption {
	var option EventRequestOption
	for i := range opts {
		opts[i](&option)
	}
	return &option
}

// NewEventRequest ...
// @param EventKey
// @param ...EventRequestOptions
// @return *EventRequest
func NewEventRequest(key EventKey, opts ...EventRequestOptions) *EventRequest {
	opt := parseERO(opts)
	return &EventRequest{
		EventKey:           key,
		EventRequestOption: opt,
	}
}
