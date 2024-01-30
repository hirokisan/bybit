package bybit

import (
	"io"
	"log"
)

var logger *log.Logger = newNoopLogger()

func SetLogger(l *log.Logger) {
	if l != nil {
		// Use provided logger.
		logger = l
	} else {
		// Disable logging.
		logger = newNoopLogger()
	}
}

func newNoopLogger() *log.Logger {
	return log.New(io.Discard, "", log.LstdFlags)
}
