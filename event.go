package of

type EventFunc = func(r EventRequester) error
type EventCallbackFunc = func(res *EventResult)

type EventTrigger interface {
	EventTrigger(event EventKey, opts ...EventRequestOptions)
}

type EventHandler interface {
	EventHandler(event EventKey, fn EventFunc) error
}

type EventListener interface {
	EventKeyAble
	EventTrigger
	EventHandler
}

type Event interface {
	EventKeyAble
	RegisterModule(name Name) (EventListener, error)
	ModuleEvent(name Name) (EventListener, bool)
}

type EventRegister interface {
	RegisterEvent(Event) error
}

type EventRequester interface {
	EventKey() EventKey
	EventRequestOption
}

type EventResult struct {
	Result any
	Error  error
}
