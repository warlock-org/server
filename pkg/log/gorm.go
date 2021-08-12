package log

import (
	"context"
	"time"

	gorm_logger "gorm.io/gorm/logger"
)

// GormLogger gorm logger
func GormLogger() gorm_logger.Interface {
	return &gormLogger{}
}

type gormLogger struct{}

func (log *gormLogger) LogMode(gorm_logger.LogLevel) gorm_logger.Interface {
	return log
}

func (*gormLogger) Info(_ context.Context, msg string, args ...interface{}) {
}

func (*gormLogger) Warn(_ context.Context, msg string, args ...interface{}) {
}

func (*gormLogger) Error(_ context.Context, msg string, args ...interface{}) {
}

func (*gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}
