package content

import (
	"io"
	"log"
	"os"
)

var debug = false
var logger = log.New(io.Discard, "", 0)

func Debug() {
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("[Content]")
	logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	debug = true
}
