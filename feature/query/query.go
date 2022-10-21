package query

import (
	"context"
)

type DataQuery interface {
	DataQuery(clientType ClientType, p uint32) (DataQueryHandler, error)
}

// DataQueryHandler is a query handler interface
type DataQueryHandler interface {
	Query(ctx context.Context, limit int, last int64) ([]byte, error)
	Last() int64
}
