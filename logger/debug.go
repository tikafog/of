package logger

import (
	"io"
	"log"
	"os"
)

var WipeMax = 256

func Discard() *log.Logger {
	return log.New(io.Discard, "", 0)
}

func Debug(name string) *log.Logger {
	return log.New(os.Stderr, name, log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
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
