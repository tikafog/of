package of

type TypeEventFunc = func(r *EventRequest) error

type EventRequest struct {
	From Name
	Type string
	Data []byte
}
