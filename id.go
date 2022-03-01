package of

type ID string

func (id ID) String() string {
	return string(id)
}

func (id ID) IsEmpty() bool {
	return id == ""
}
