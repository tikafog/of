package of

import (
	"context"

	"github.com/tikafog/of/buffers/content"
)

// Core
// @Description: the BasicCore of the framework
type Core interface {
	Context() context.Context
	State() State
	RegisterMessageHandler(protocol Protocol, handler MessageHandler) error
	QueryHandler(ct content.Type) (DataQueryHandler, error)

	CoreFeature
}
