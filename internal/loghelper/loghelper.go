package loghelper

import (
	"github.com/sirupsen/logrus"
	"time"
)

// Event stores messages to log later, from our standard interface.
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats.
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger.
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// LogCalls is a standard error message.
func (l *StandardLogger) LogCalls(uri, method string, duration time.Duration) {
	l.WithFields(logrus.Fields{
		"uri":      uri,
		"method":   method,
		"duration": duration,
	}).Info("Call logs")
}

// LogPanic is a standard error message.
func (l *StandardLogger) LogPanic(err interface{}) {
	l.WithFields(logrus.Fields{
		"err":      err,
	}).Warn("There is some error, recovered")
}

// LogError is a standard error message.
func (l *StandardLogger) LogError(uri, method string, err interface{}) {
	logrus.WithFields(logrus.Fields{
		"uri":      uri,
		"method":   method,
		"err": err,
	}).Fatal("There is some error, recovered")
}
