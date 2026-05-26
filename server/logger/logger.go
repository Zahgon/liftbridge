package logger

import (
	"io"
	"sync"

	gnatsd "github.com/nats-io/nats-server/v2/server"
	log "github.com/sirupsen/logrus"
)

// Logger interface is used to allow tests to inject custom loggers.
type Logger interface {
	Fatalf(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Fatal(...interface{})
	Silent(bool)
	Prefix(string)
}

type logger struct {
	*log.Logger
	oldOut io.Writer
	prefix string
	mu     sync.RWMutex
}

// NewLogger returns a new Logger instance backed by Logrus.
func NewLogger(level uint32) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// Fatalf logs a fatal error.
func (l *logger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf logs a debug statement.
func (l *logger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf logs an error statement.
func (l *logger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Infof logs an info statement.
func (l *logger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf logs an warn statement.
func (l *logger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Debug logs a debug statement.
func (l *logger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

// Warn logs a warn statement.
func (l *logger) Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

// Info logs an info statement.
func (l *logger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

// Fatal logs a fatal error.
func (l *logger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

// Silent is used to enable and disable log silencing. Silent must be called
// with true before it can be called with false.
func (l *logger) Silent(enable bool) { _ = "STUB: not implemented"; return }

// Prefix all log output with the given string. Pass an empty string to clear
// any previously set prefix.
func (l *logger) Prefix(prefix string) { _ = "STUB: not implemented"; return }

func (l *logger) prefixFormat(format string) string { _ = "STUB: not implemented"; return "" }

func (l *logger) prefixVars(v []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// natsLogger implements the NATS server logger interface by writing log
// messages to a Liftbridge logger.
type natsLogger struct {
	logger Logger
}

// NewNATSLogger creates a NATS logger that writes log messages to the given
// Logger.
func NewNATSLogger(logger Logger, enabled bool) gnatsd.Logger {
	_ = "STUB: not implemented"
	return *new(gnatsd.Logger)
}

// Noticef logs a notice statement.
func (n *natsLogger) Noticef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf logs a warning statement.
func (n *natsLogger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf logs a fatal error.
func (n *natsLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf logs an error.
func (n *natsLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf logs a debug statement.
func (n *natsLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Tracef logs a trace statement.
func (n *natsLogger) Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// noopNATSLogger implements the NATS server logger interface by performing a
// no-op for each log statement with the exception of Fatalf.
type noopNATSLogger struct {
	logger Logger
}

// Noticef is a no-op.
func (n *noopNATSLogger) Noticef(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	// No-op

	// Warnf is a no-op.
	return
}

func (n *noopNATSLogger) Warnf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	// No-op

	// Fatalf logs a fatal error.
	return
}

func (n *noopNATSLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf is a no-op.
func (n *noopNATSLogger) Errorf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	// No-op

	// Debugf is a no-op.
	return
}

func (n *noopNATSLogger) Debugf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	// No-op

	// Tracef is a no-op.
	return
}

func (n *noopNATSLogger) Tracef(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	// No-op
	return
}
