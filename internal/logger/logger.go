package logger

import (
    "os"
    
    "github.com/sirupsen/logrus"
)

var log *logrus.Logger

func Init(level string) {
    log = logrus.New()
    log.SetOutput(os.Stdout)
    log.SetFormatter(&logrus.JSONFormatter{})
    
    logLevel, err := logrus.ParseLevel(level)
    if err != nil {
        logLevel = logrus.InfoLevel
    }
    log.SetLevel(logLevel)
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

func GetLogger() *logrus.Logger {
    return log
}
