package of

type Node interface {
	Connection
	Streamer
	Inquirer

	ID() ID
	HandleMessage(protocol Protocol, handler MessageHandler) error
}
