package of

type EventKey struct {
	_from     Name
	_receiver Name
	_type     string
}

func (e *EventKey) Type() string {
	return e._type
}

func (e *EventKey) From() Name {
	return e._from
}

func (e *EventKey) Receiver() Name {
	return e._receiver
}

func NewEventKey(from, receiver Name, p string) *EventKey {
	return &EventKey{
		_from:     from,
		_receiver: receiver,
		_type:     p,
	}
}
