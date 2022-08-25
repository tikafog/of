package of

import (
	"github.com/tikafog/of/buffers/content"
)

// Core
// @Description: the BasicCore of the framework
type Core interface {
	//Context() context.Context
	State() State

	QueryHandler(ct content.Type) (DataQueryHandler, error)

	CoreFeature
}
