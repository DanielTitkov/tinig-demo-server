package prepare

import "go.uber.org/zap"

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(msg, info string) {
	l.logger.Info(msg,
		zap.String("info", info),
	)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg,
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Error(msg,
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Sync() {
	l.logger.Sync()
}
