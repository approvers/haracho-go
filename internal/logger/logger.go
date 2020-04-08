package logger

type Logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})

	Infofln(string, ...interface{})
	Errorfln(string, ...interface{})
	Debugfln(string, ...interface{})
}
