package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

// Init initializes logger
func Init(app string, lvl string) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l, err := zerolog.ParseLevel(lvl)
	if err != nil {
		return err
	}
	logger = zerolog.New(os.Stderr).With().Timestamp().Str("app", app).Logger().Level(l)

	return nil
}

// Info returns zerolog event
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn returns zerolog event
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error returns zerolog event
func Error() *zerolog.Event {
	return logger.Error()
}

// Debug returns zerolog event
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Fatal returns zerolog event
func Fatal() *zerolog.Event {
	return logger.Fatal()
}

// Panic returns zerolog event
func Panic() *zerolog.Event {
	return logger.Panic()
}

// Trace returns zerolog event
func Trace() *zerolog.Event {
	return logger.Trace()
}

// FatalOnError logs error and exits
func FatalOnError(msg string, err error) {
	if err != nil {
		logger.Fatal().Err(err).Msg(msg)
	}
}

// Err logs error
func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

// Printf prints formatted message
func Printf(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
}
