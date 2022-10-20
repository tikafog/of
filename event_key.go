package of

type EventKey interface {
	KeyID() uint64
	Type() string
	From() Name
	Receiver() Name
}
