package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.Out = os.Stdout
}

// SetLogLevel sets log level
func SetLogLevel(level logrus.Level) {
	log.Level = level
}

// SetLogFormatter sets log formatter
func SetLogFormatter(formatter logrus.Formatter) {
	log.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if log.Level >= logrus.DebugLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args)
	}
}

// DebugWithFields logs a message with fields at level Debug on the standard logger.
func DebugWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.DebugLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if log.Level >= logrus.InfoLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// InfoWithFields logs a message with fields at level Info on the standard logger.
func InfoWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.InfoLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if log.Level >= logrus.WarnLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields logs a message at level Warn on the standard logger.
func WarnWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.WarnLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if log.Level >= logrus.ErrorLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// ErrorWithFields logs a message at level Error on the standard logger.
func ErrorWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.ErrorLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if log.Level >= logrus.FatalLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// FatalWithFields logs a message with fields at level Fatal on the standard logger.
func FatalWithFields(l interface{}, f Fields) {
	if log.Level >= logrus.FatalLevel {
		entry := log.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if log.Level >= logrus.PanicLevel {
		entry := log.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// PanicWithFields logs a message with fields at level Panic on the standard logger.
func PanicWithFields(l interface{}, f Fields) {
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
