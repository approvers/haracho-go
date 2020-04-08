package logger

type Logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
}
