package query

import (
	"context"

	"github.com/tikafog/of/buffers/content"
)

type DataQuery interface {
	DataQuery(clientType ClientType, p uint32) (DataQueryHandler, error)
}

// DataQueryHandler is a query handler interface
type DataQueryHandler interface {
	Query(ctx context.Context, limit int, last int64) ([]byte, error)
	Last() int64
}

// ContentClientType ...
// @param content.Type
// @return ClientType
func ContentClientType(tp content.Type) ClientType {
	switch tp {
	case content.TypeUpdate:
		return ClientTypeUpgrade
	case content.TypeBootstrap:
		return ClientTypeBootnode
	}
	return ClientTypeMedia
}
