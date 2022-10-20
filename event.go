package of

type EventFunc = func(r EventRequester) error
type EventCallbackFunc = func(res *EventResult)

//type EventTrigger interface {
//	Trigger(key EventKey, opts ...EventRequestOptions) error
//}

type KeyMaker interface {
	To(receiver Name, p string) EventKey
}

type Event interface {
	KeyMaker(from Name) KeyMaker
	Register(key EventKey, fn EventFunc) error
	Trigger(key EventKey, opts ...EventRequestOptions) error
}

type EventRegister interface {
	RegisterEvent(Event) error
}

type EventRequester interface {
	EventKey
	EventRequestOption
}

type EventResult struct {
	Result any
	Error  error
}
