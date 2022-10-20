package of

type EventKey interface {
	Type() string
	From() Name
	Receiver() Name
}
