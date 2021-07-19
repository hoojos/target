package log

import "github.com/sirupsen/logrus"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05.999",
	})
	DefaultLogger = logrus.NewEntry(logrus.StandardLogger())
}

var DefaultLogger logrus.FieldLogger

func WithFields(fields logrus.Fields) *logrus.Entry {
	return DefaultLogger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return DefaultLogger.WithError(err)
}

func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

func Warning(args ...interface{}) {
	DefaultLogger.Warning(args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}
