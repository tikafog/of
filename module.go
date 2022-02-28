package of

import (
	"context"
)

//Name returns the names of all the modules
//ENUM(admin,center,node,bootNode,instruct,max)
type Name int

type Module interface {
	Valid() bool
	Run(ctx context.Context) error
	Destroy() error
	Name() Name

	//this all calls before run
	PreinstallCore(core Core) error
	PreinstallAPI(api API) error
}
