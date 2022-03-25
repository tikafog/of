package of

import (
	"encoding/json"
)

type TypeHandleFunc = func(conn Conn, data json.RawMessage, args ...Arg) error

type DataHandler interface {
	DataHandle(conn Conn, data json.RawMessage, args ...Arg) error
	Query(limit int, last int64) ([]byte, error)
	Last() int64
}
