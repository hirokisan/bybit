package bybit

import (
	"io"
	"log"
	"os"
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

func newDefaultLogger() *log.Logger {
	return log.New(os.Stderr, "Bybit-golang", log.LstdFlags)
}

func newNoopLogger() *log.Logger {
	return log.New(io.Discard, "", log.LstdFlags)
}
