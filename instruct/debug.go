package instruct

import (
	"github.com/tikafog/of/logger"
)

var debug = false
var log = logger.Discard()
var Wipe = logger.Wipe

func Debug() {
	log = logger.Debug("[Instruct]")
}
