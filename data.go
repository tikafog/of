package of

import (
	"context"
)

// DataHandler is a data handler interface\
type DataHandler[T any, A any] interface {
	DataUpdateHandler[T, A]
	DataQueryHandler[A]
}

//DataQueryHandler is a query handler interface
type DataQueryHandler[A any] interface {
	Query(ctx context.Context, limit int, last int64, args ...A) ([]byte, error)
	Last(args ...A) int64
}

// DataUpdateHandler is a update handler interface
type DataUpdateHandler[T any, A any] interface {
	Update(ctx context.Context, data T, args ...A) (int, error)
}
