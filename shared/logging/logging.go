package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

type Fields map[string]interface{}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.Formatter = &log.TextFormatter{}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.Out = os.Stdout

	// Only log the warning severity or above.
	logger.Level = log.DebugLevel
}

func WithFields(fields Fields) *log.Entry {
	return logger.WithFields(log.Fields(fields))
}
