package of

type Node interface {
	Connection
	Streamer
	Inquirer
	Answer

	ID() ID
	HandleMessage(protocol Protocol, handler MessageHandler) error
}
