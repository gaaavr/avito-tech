package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Log - event logging object
type Log struct {
	*logrus.Entry
}

// New - constructor function for Log// User - Interface describing the user entity
func New() *Log {
	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{
		DisableColors: false,
	}
	l.SetLevel(logrus.DebugLevel)
	l.SetOutput(os.Stdout)
	return &Log{logrus.NewEntry(l)}
}
