package of

import (
	"context"
)

// DataHandler is a data handler interface\
type DataHandler[T any, A any] interface {
	DataUpdateHandler[T, A]
	DataQueryHandler
}

//DataQueryHandler is a query handler interface
type DataQueryHandler interface {
	Query(ctx context.Context, limit int, last int64) ([]byte, error)
	Last() int64
}

// DataUpdateHandler is a update handler interface
type DataUpdateHandler[T any, A any] interface {
	Update(ctx context.Context, data T, args ...A) (int, error)
}
