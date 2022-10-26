package of

type Node interface {
	ID() ID
	IsAdmin() bool

	Connection
	Streamer
	Inquirer
}
