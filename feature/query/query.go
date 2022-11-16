package query

import (
	"context"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/feature/source"
)

type DataQuery interface {
	DataQuery(clientType source.Type, p uint32) (DataQueryHandler, error)
}

// DataQueryHandler is a query handler interface
type DataQueryHandler interface {
	Query(ctx context.Context, limit int, last int64) ([]byte, error)
	Last() int64
}

// ContentSourceType ...
// @param content.Type
// @return Type
func ContentSourceType(tp content.Type) source.Type {
	switch tp {
	case content.TypeUpdate:
		return source.TypeUpgrade
	case content.TypeBootstrap:
		return source.TypeBootnode
	}
	return source.TypeMedia
}
