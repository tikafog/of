package of

type Tools interface {
	Identifier() Identifier
}

type Identifier interface {
	Decode(id string) (ID, error)
	Encode(id ID) string
}
