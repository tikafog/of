package of

import (
	"context"
)

//Name returns the names of all the modules
type Name string

type Module interface {
	Valid() bool
	Init() error
	Run(ctx context.Context) error
	Destroy()
	Name() Name

	//this all calls before run
	PreloadCore(core Core) error

	//this all calls after run
	RegisterAPI(api API) error
}
