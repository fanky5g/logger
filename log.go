package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields
type Logger struct {
	*logrus.Logger
}

var (
	log  *Logger
	hook *logrusly.LogglyHook
)

func New(logglyToken, logFilePath string) *Logger {
	log = &Logger{
		Logger: logrus.New(),
	}

	if logglyToken != "" {
		hook = logrusly.NewLogglyHook(logglyToken, "https://logs-01.loggly.com/bulk/", logrus.WarnLevel|logrus.DebugLevel|logrus.ErrorLevel|logrus.FatalLevel|logrus.InfoLevel, "go", "logrus")
		log.Hooks.Add(hook)
	}

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open log file", os.Stdout, ":", err)
		}

		log.Out = io.Writer(file)
	}

	return log
}

// SetLogLevel sets log level
func (log *Logger) SetLogLevel(level logrus.Level) {
	log.Level = level
}

// DebugMode sets logger in debug mode
func (log *Logger) DebugMode() *Logger {
	log.SetLogLevel(logrus.DebugLevel)
	return log
}

// InfoMode sets logger in info mode
func (log *Logger) InfoMode() *Logger {
	log.SetLogLevel(logrus.InfoLevel)
	return log
}

// ErrorMode sets logger in error mode
func (log *Logger) ErrorMode() *Logger {
	log.SetLogLevel(logrus.ErrorLevel)
	return log
}

// FatalMode sets logger in fatal mode
func (log *Logger) FatalMode() *Logger {
	log.SetLogLevel(logrus.FatalLevel)
	return log
}

// SetLogFormatter sets log formatter
func (log *Logger) SetLogFormatter(formatter logrus.Formatter) {
	log.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logger.
func (log *Logger) Debug(args ...interface{}) {
	if log.Level >= logrus.DebugLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args)
		if hook != nil {
			hook.Flush()
		}
	}
}

// DebugWithFields logs a message with fields at level Debug on the standard logger.
func (log *Logger) DebugWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.DebugLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
		if hook != nil {
			hook.Flush()
		}
	}
}

// Info logs a message at level Info on the standard logger.
func (log *Logger) Info(args ...interface{}) {
	if log.Level >= logrus.InfoLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
		if hook != nil {
			hook.Flush()
		}
	}
}

// InfoWithFields logs a message with fields at level Info on the standard logger.
func (log *Logger) InfoWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.InfoLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Info(l)
		if hook != nil {
			hook.Flush()
		}
	}
}

// Warn logs a message at level Warn on the standard logger.
func (log *Logger) Warn(args ...interface{}) {
	if log.Level >= logrus.WarnLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
		if hook != nil {
			hook.Flush()
		}
	}
}

// WarnWithFields logs a message at level Warn on the standard logger.
func (log *Logger) WarnWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.WarnLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
		if hook != nil {
			hook.Flush()
		}
	}
}

// Error logs a message at level Error on the standard logger.
func (log *Logger) Error(args ...interface{}) {
	if log.Level >= logrus.ErrorLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
		if hook != nil {
			hook.Flush()
		}
	}
}

// ErrorWithFields logs a message at level Error on the standard logger.
func (log *Logger) ErrorWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.ErrorLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
		if hook != nil {
			hook.Flush()
		}
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func (log *Logger) Fatal(args ...interface{}) {
	if log.Level >= logrus.FatalLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
		if hook != nil {
			hook.Flush()
		}
	}
}

// FatalWithFields logs a message with fields at level Fatal on the standard logger.
func (log *Logger) FatalWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.FatalLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
		if hook != nil {
			hook.Flush()
		}
	}
}

// Panic logs a message at level Panic on the standard logger.
func (log *Logger) Panic(args ...interface{}) {
	if log.Level >= logrus.PanicLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// PanicWithFields logs a message with fields at level Panic on the standard logger.
func (log *Logger) PanicWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.PanicLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
