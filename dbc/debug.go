package dbc

import (
	"github.com/tikafog/of/logger"
)

var log = logger.WithGroup("[DBC]")

func mediaDebug(args ...interface{}) {
	log.Debug("media database debug", args...)
}

func kernelDebug(args ...interface{}) {
	log.Debug("kernel database debug", args...)
}
func upgradeDebug(args ...interface{}) {
	log.Debug("upgrade database debug", args...)
}
func bootnodeDebug(args ...interface{}) {
	log.Debug("bootnode database debug", args...)
}
