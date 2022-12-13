package of

type EventFunc = func(r EventRequester) error
type EventCallbackFunc = func(res *EventResult)

type KeyMaker interface {
	To(receiver Name, p string) EventKey
}

type EventTrigger interface {
	EventTrigger(event string, opts ...EventRequestOptions)
}

type EventHandler interface {
	EventHandleEvent(event string, fn EventFunc) error
}

type EventListener interface {
	EventTrigger
	EventHandler
}

type Event interface {
	RegisterModule(name Name) (EventListener, error)
	ModuleEvent(name) (EventListener, bool)
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
