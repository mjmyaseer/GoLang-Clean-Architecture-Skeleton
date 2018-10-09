package logger

import (
	"bytes"
	"context"
	"go-sample/utils/go-util/log"
)

type LogEntry struct {
	message  bytes.Buffer
	module   string
	function string
	query    string
}

var (
	ERROR = `ERROR`
	FATAL = `FATAL`
	INFO  = `INFO`
	DEBUG = `DEBUG`
	TRACE = `TRACE`
)

func (entry LogEntry) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	log.ErrorContext(ctx, message, params)
}

func (entry LogEntry) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	log.InfoContext(ctx, message, params)
}

func (entry LogEntry) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	log.DebugContext(ctx, message, params)

}

func (entry LogEntry) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	log.FatalContext(ctx, message, params)
}

//Log to remote service
func (entry LogEntry) logRemote() {
	// TODO implement this
}

func Log() LogEntry {
	lg := LogEntry{}
	return lg
}
