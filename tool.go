package of

type Tools interface {
	Identifier() Identifier
	Net() Net
}

type Identifier interface {
	Decode(id string) (ID, error)
	Encode(id ID) string
}

type Addr interface {
	ID() ID
	Addrs() []string
	AddrStrings() []string
}

type Net interface {
	ParseStringAddr(addr string) (Addr, error)
	ParseStringAddrs(addrs ...string) ([]Addr, error)
	StringAddr(addr Addr) ([]string, error)
}
