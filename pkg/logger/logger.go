package logger

type LogFormat interface {
	Info(msg string)
	Errorf(msg string, err error)
	Error(err error)
	Warning(err error)
	Warningf(msg string, err error)
}
