package of

type TypeEventFunc = func(r *EventRequest) error

type EventRequest struct {
	Receiver Name
	Type     string
	Value    interface{}
}
