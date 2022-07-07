package of

import (
	"context"
)

//type TypeHandleFunc = func(conn Conn, data json.RawMessage, args ...Arg) error

// DataHandler
// @Description: DataHandler
type DataHandler[T any, V any] interface {
	DataUpdateHandler[T, V]
	DataQueryHandler
}

type DataQueryHandler interface {
	Query(ctx context.Context, limit, offset int64) ([]byte, error)
	Last(ctx context.Context) int64
}

type DataUpdateHandler[T any, V any] interface {
	Update(ctx context.Context, data T, args ...V) error
}
