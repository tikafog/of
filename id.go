package of

type ID interface {
	IsEmpty() bool
	String() string
	HexString() string
}
