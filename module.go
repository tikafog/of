package of

import (
	"context"
	"strings"
)

//Name returns the names of all the modules
type Name string

type ModuleStarter interface {
	Init() error
	Run(ctx context.Context) error
	Destroy()

	//this all calls before run
	PreloadCore(core Core) error

	//this all calls after run
	RegisterAPI(api API) error

	Module
}

type ModuleStater interface {
	Name() Name
	IsNil() bool
	IsRunning() bool
}

type Module interface {
	ModuleStater
	Data(limit int, last int64) ([]byte, error)
	Wait()
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
