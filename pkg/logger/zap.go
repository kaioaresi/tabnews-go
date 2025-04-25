package logger

import "go.uber.org/zap"

type Logger struct {
	Sugar *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	return &Logger{
		Sugar: logger.Sugar(),
	}, err
}

func (l *Logger) Info(msg string) {
	l.Sugar.Infof("%s", msg)
}

func (l *Logger) Errorf(msg string, err error) {
	l.Sugar.Errorf("%s - %s", msg, err)
}

func (l *Logger) Error(err error) {
	l.Sugar.Error(err)
}

func (l *Logger) Warning(err error) {
	l.Sugar.Warn(err)
}

func (l *Logger) Warningf(msg string, err error) {
	l.Sugar.Warnf("%v - %v", msg, err)
}
