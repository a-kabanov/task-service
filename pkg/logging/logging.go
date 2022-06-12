package logging

import (
	"io"
	logging "log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

const log_path string = "logs/task-info.log"

var (
	log       *logrus.Logger
	FBuchName string
)

var lock = sync.Mutex{}

func init() {
	mainLogFile, err := os.OpenFile(log_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	log = logrus.New()
	//log.Formatter = &logrus.JSONFormatter{}

	log.SetReportCaller(false)

	mw := io.MultiWriter(os.Stdout, mainLogFile)
	log.SetOutput(mw)

}

// Info ...
func Info(format string, v ...interface{}) {
	lock.Lock()
	log.Infof(format, v...)
	lock.Unlock()
}

// Warn ...
func Warn(format string, v ...interface{}) {
	lock.Lock()
	log.Warnf(format, v...)
	lock.Unlock()
}

// Error ...
func Error(format string, v ...interface{}) {
	lock.Lock()
	log.Errorf(format, v...)
	lock.Unlock()
}

// Fatal ...
func Fatal(format string, v ...interface{}) {
	lock.Lock()
	logging.Fatalf(format, v...)
	lock.Unlock()
}
