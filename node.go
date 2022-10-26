package of

type Node interface {
	ID() ID

	Connection
	Streamer
	Inquirer
}
