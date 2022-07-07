package of

import (
	"context"
)

//type TypeHandleFunc = func(conn Conn, data json.RawMessage, args ...Arg) error

// DataHandler
// @Description: DataHandler
type DataHandler[T any, V any] interface {
	DataUpdateHandler[T, V]
	DataQueryHandler[T, V]
}

type DataQueryHandler[T any, V any] interface {
	Query(ctx context.Context, data T, args ...V) ([]byte, error)
	Last(ctx context.Context, data T, args ...V) int64
}

type DataUpdateHandler[T any, V any] interface {
	Update(ctx context.Context, data T, args ...V) error
}
