package of

import (
	"context"
	"strings"
)

//Name returns the names of all the modules
type Name string

type Module interface {
	Valid() bool
	Init() error
	Run(ctx context.Context) error
	IsRunning() bool
	Destroy()
	Name() Name

	//this all calls before run
	PreloadCore(core Core) error

	//this all calls after run
	RegisterAPI(api API) error
}

func (n Name) String() string {
	return strings.ToLower(string(n))
}

func (n Name) Compare(other Name) bool {
	return n.String() == other.String()
}

func CompareName(n, o Name) bool {
	return n.String() == o.String()
}
