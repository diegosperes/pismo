package util

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormL "gorm.io/gorm/logger"
)

func getGormLogLevel(levelName string) gormL.LogLevel {
	switch strings.ToUpper(levelName) {
	case "ERROR":
		return gormL.Error
	case "WARN":
		return gormL.Warn
	case "INFO":
		return gormL.Info
	case "DEBUG":
		return gormL.Info
	default:
		return gormL.Silent
	}
}

type gormWritter struct{}

func (w gormWritter) Printf(msg string, data ...interface{}) {
	msg = fmt.Sprintf(msg, data...)
	slog.Info(msg)
}

type gormLogger struct {
	logger gormL.Interface
}

func newGormLogger(logLevelName string) *gormLogger {
	return &gormLogger{
		logger: gormL.New(
			&gormWritter{},
			gormL.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  getGormLogLevel(logLevelName),
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	}
}

func (l gormLogger) LogMode(level gormL.LogLevel) gormL.Interface {
	return l
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	slog.InfoContext(ctx, fmt.Sprintf(msg, data...))
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	slog.WarnContext(ctx, fmt.Sprintf(msg, data...))
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	slog.ErrorContext(ctx, fmt.Sprintf(msg, data...))
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	l.logger.Trace(ctx, begin, fc, err)
}

func SetupDatabaseConnection(dbSettings DatabaseSettings, logLevelName string) (*gorm.DB, error) {
	if len(dbSettings.LogLevelName) > 0 {
		logLevelName = dbSettings.LogLevelName
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbSettings.Host,
		dbSettings.User,
		dbSettings.Pass,
		dbSettings.Name,
		dbSettings.Port,
		dbSettings.SSLMode,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger:         newGormLogger(logLevelName),
			TranslateError: true,
		},
	)

	return db, err
}
