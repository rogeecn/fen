package fen

import "github.com/sirupsen/logrus"

var (
	AlwaysStatusOK = false
	MessageSuccess = "OK"
	DebugMode      = false
)

var logger Logger = logrus.New()

func SetLogger(l Logger) {
	logger = l
}

func SetDebug(b bool) {
	DebugMode = b
}

type Logger interface {
	Error(args ...interface{})
}
