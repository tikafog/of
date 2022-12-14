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
	EventTrigger
	EventHandler
}

type Event interface {
	Key(str string) EventKey
	KeyName(EventKey) string
	RegisterModule(name Name) (EventListener, error)
	ModuleEvent(name) (EventListener, bool)
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
