package observability

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Fatal(format string, args ...interface{})
}

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogger() Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return &LogrusLogger{
		logger: log,
	}
}

func (l *LogrusLogger) Info(msg string, fields ...map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields[0]).Info(msg)
	} else {
		l.logger.Info(msg)
	}
}

func (l *LogrusLogger) Error(msg string, fields ...map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields[0]).Error(msg)
	} else {
		l.logger.Error(msg)
	}
}

func (l *LogrusLogger) Warn(msg string, fields ...map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields[0]).Warn(msg)
	} else {
		l.logger.Warn(msg)
	}
}

func (l *LogrusLogger) Debug(msg string, fields ...map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields[0]).Debug(msg)
	} else {
		l.logger.Debug(msg)
	}
}

func (l *LogrusLogger) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
