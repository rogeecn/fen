package fen

import "github.com/sirupsen/logrus"

var (
	AlwaysStatusOK = false
	MessageSuccess = "OK"
	DebugMode      = false
)

var LOG Logger = logrus.New()

type Logger interface {
	Error(args ...interface{})
}
