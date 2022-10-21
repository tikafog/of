package of

import (
	"context"

	"github.com/tikafog/of/dbc"
)

// Core
// @Description: the BasicCore of the framework
type Core interface {
	Context() context.Context
	State() State

	DataQuery(clientType dbc.ClientType, p uint32) (DataQueryHandler, error)

	CoreFeature
}
