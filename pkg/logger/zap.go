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
	defer l.Sugar.Sync()
}

func (l *Logger) Errorf(msg string, err error) {
	l.Sugar.Errorf("%s - %s", msg, err)
	defer l.Sugar.Sync()
}

func (l *Logger) Error(err error) {
	l.Sugar.Error(err)
	defer l.Sugar.Sync()
}

func (l *Logger) Warning(err error) {
	l.Sugar.Warn(err)
	defer l.Sugar.Sync()
}

func (l *Logger) Warningf(msg string, err error) {
	l.Sugar.Warnf("%v - %v", msg, err)
	defer l.Sugar.Sync()
}
