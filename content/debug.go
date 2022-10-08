package content

import (
	"io"
	"log"
	"os"
)

var debug = false
var logger = log.New(io.Discard, "", 0)
var WipeMax = 256

func Debug() {
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("[Content]")
	logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	debug = true
}

func WipeString(str string) string {
	if len(str) < WipeMax {
		return str
	}
	return str[:WipeMax]
}

func Wipe(bytes []byte) string {
	if len(bytes) < WipeMax {
		return string(bytes)
	}
	return string(bytes[:WipeMax])
}
