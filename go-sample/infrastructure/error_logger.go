package infrastructure

import (
	"go-sample/utils/go-util/log"
)

type ErrorLogger struct {
}

func (*ErrorLogger) Log(keyvals ...interface{}) error {
	log.Error(log.WithPrefix(`infrastructure.go-kit.errors`, keyvals))
	return nil
}
