package of

import (
	"context"

	"github.com/tikafog/of/feature/query"
)

// Core
// @Description: the BasicCore of the framework
type Core interface {
	Context() context.Context
	State() State

	query.DataQuery
	CoreFeature
}
