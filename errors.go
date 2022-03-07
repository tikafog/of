package of

import (
	"errors"
)

var ErrNilCore = errors.New("core is nil")
var ErrFatherNotFound = errors.New("father not found")
var ErrNoDataFound = errors.New("no data found")
var ErrChannelClosed = errors.New("channel closed")
