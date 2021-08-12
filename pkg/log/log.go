package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = zap.NewExample()

// SetLevel set level
func SetLevel(level zapcore.Level) {
	logger.WithOptions(zap.IncreaseLevel(level))
}

// Info info
func Info(args ...interface{}) {
	logger.Sugar().Info(args)
}

// Error error
func Error(args ...interface{}) {
	logger.Sugar().Error(args)
}

// Warn warn
func Warn(args ...interface{}) {
	logger.Sugar().Warn(args)
}

// Debug debug
func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

// Infof Infof
func Infof(msg string, args ...interface{}) {
	logger.Sugar().Infof(msg, args)
}

// Errorf Errorf
func Errorf(msg string, args ...interface{}) {
	logger.Sugar().Errorf(msg, args)
}

// Warnf Warnf
func Warnf(msg string, args ...interface{}) {
	logger.Sugar().Warnf(msg, args)
}

// Debugf Debugf
func Debugf(msg string, args ...interface{}) {
	logger.Sugar().Debugf(msg, args...)
}
