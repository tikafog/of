package of

//type TypeHandleFunc = func(conn Conn, data json.RawMessage, args ...Arg) error

// DataHandler
// @Description: DataHandler
type DataHandler[T any, V any] interface {
	DataHandle(data T, args ...V) error
	Query(limit int, last int64) ([]byte, error)
	Last() int64
}
