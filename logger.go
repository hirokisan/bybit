package bybit

import (
	"io"
	"log"
)

type Logger interface {
	Println(values ...any)
}

var logger Logger = newNoopLogger()

func SetLogger(l Logger) {
	if l != nil {
		// Use provided logger.
		logger = l
	} else {
		// Disable logging.
		logger = newNoopLogger()
	}
}

func newNoopLogger() Logger {
	return log.New(io.Discard, "", log.LstdFlags)
}
