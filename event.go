package of

type EventFunc = func(r EventRequester) error
type EventCallbackFunc = func(res *EventResult)

type EventTrigger interface {
	EventKeyAble
	EventTrigger(event EventKey, opts ...EventRequestOptions)
}

type EventHandler interface {
	EventKeyAble
	EventHandler(event EventKey, fn EventFunc) error
}

type EventListener interface {
	EventTrigger
	EventHandler
}

type Event interface {
	EventKeyAble
	RegisterModule(name Name) (EventListener, error)
	ModuleEvent(name Name) (EventListener, bool)
}

type EventRequester interface {
	EventKey() EventKey
	EventRequestOption
}

type EventResult struct {
	Result any
	Error  error
}
