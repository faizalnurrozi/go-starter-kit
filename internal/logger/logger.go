package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func Init(level string) {
	log = logrus.New()

	if level == "silent" {
		log.SetOutput(io.Discard)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	log.SetLevel(logLevel)
}

func SetOutput(w io.Writer) {
	if log != nil {
		log.SetOutput(w)
	}
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}

func InfoStructured(event string, routing string, message any) {
	log.WithFields(logrus.Fields{
		"event":   event,
		"routing": routing,
		"message": message,
		"pid":     os.Getpid(),
	}).Info("Ini message")
}

func GetLogger() *logrus.Logger {
	return log
}
