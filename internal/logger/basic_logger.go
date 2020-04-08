package logger

import (
	"log"
)

type BasicLogger struct {
	Logger *log.Logger
}

func (bl *BasicLogger) Infof(format string, v ...interface{}) {
	bl.Logger.Printf("[Info] "+format, v)
}

func (bl *BasicLogger) Errorf(format string, v ...interface{}) {
	bl.Logger.Printf("[Error] "+format, v)
}

func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	bl.Logger.Printf("[Debug] "+format, v)
}

func (bl *BasicLogger) Infofln(format string, v ...interface{}) {
	bl.Logger.Printf("[Info] "+format+"\n", v)
}

func (bl *BasicLogger) Errorfln(format string, v ...interface{}) {
	bl.Logger.Printf("[Error] "+format+"\n", v)
}

func (bl *BasicLogger) Debugfln(format string, v ...interface{}) {
	bl.Logger.Printf("[Debug] "+format+"\n", v)
}
