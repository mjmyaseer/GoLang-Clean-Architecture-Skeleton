package log

import (
	"context"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"log"
	"os"
	"runtime"
	//"net"
)

var nativeLog *log.Logger
var errorLog *log.Logger

var (
	fatal = `FATAL`
	err   = `ERR`
	warn  = `WARN`
	info  = `INFO`
	debug = `DEBUG`
	trace = `TRACE`
)

var logColors = map[string]string{
	`FATAL`: BgRed(`[FATAL]`).String(),
	`ERROR`: BgRed(`[ERROR]`).String(),
	`WARN`:  BgBrown(`[WARN]`).String(),
	`INFO`:  BgBlue(`[INFO]`).String(),
	`DEBUG`: BgCyan(`[DEBUG]`).String(),
	`TRACE`: BgMagenta(`[TRACE]`).String(),
}

var logTypes = map[string]int{
	`FATAL`: 1,
	`ERROR`: 2,
	`WARN`:  3,
	`INFO`:  4,
	`DEBUG`: 5,
	`TRACE`: 6,
}

func init() {
	nativeLog = log.New(os.Stdout, ``, log.LstdFlags|log.Lmicroseconds)
	errorLog = log.New(os.Stderr, ``, log.LstdFlags|log.Lmicroseconds)
}

func colored(typ string) string {
	if logConfig.Colors {
		return logColors[typ]
	}

	return `[` + typ + `]`
}

//isLoggable Check whether the log type is loggable under current configurations
func isLoggable(logType string) bool {
	return logTypes[logType] <= logTypes[logConfig.Level]
}

func toString(id string, typ string, message interface{}, params ...interface{}) string {

	var messageFmt = "%s %s %v"

	return fmt.Sprintf(messageFmt,
		typ,
		fmt.Sprintf("%+v", message),
		fmt.Sprintf("%+v", params))
}

func ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(err, ctx, message, colored(`ERROR`), params...)
}

func WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(warn, ctx, message, colored(`WARN`), params...)
}

func InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(info, ctx, message, colored(`INFO`), params...)
}

func DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(debug, ctx, message, colored(`DEBUG`), params...)
}

func TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(trace, ctx, message, colored(`TRACE`), params...)
}

func Error(message interface{}, params ...interface{}) {
	logEntry(err, message, colored(`ERROR`), params...)
}

func Warn(message interface{}, params ...interface{}) {
	logEntry(warn, message, colored(`WARN`), params...)
}

func Info(message interface{}, params ...interface{}) {
	logEntry(info, message, colored(`INFO`), params...)
}

func Debug(message interface{}, params ...interface{}) {
	logEntry(debug, message, colored(`DEBUG`), params...)
}

func Trace(message interface{}, params ...interface{}) {
	logEntry(trace, message, colored(`TRACE`), params...)
}

func Fatal(message interface{}, params ...interface{}) {
	logEntry(fatal, message, colored(`FATAL`), params...)
}

func Fataln(message interface{}, params ...interface{}) {
	logEntry(fatal, message, colored(`FATAL`), params...)
}

func FatalContext(ctx context.Context, message interface{}, params interface{}) {
	logEntry(fatal, message, colored(`FATAL`), params)
}

func logEntryContext(logType string, ctx context.Context, message interface{}, color string, params ...interface{}) {
	logEntry(logType, message, color, params...)
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

func logEntry(logType string, message interface{}, color string, params ...interface{}) {

	if !isLoggable(logType) {
		return
	}

	var file string
	var line int
	if logConfig.FilePath {
		_, f, l, ok := runtime.Caller(2)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		file = f
		line = l

		message = fmt.Sprintf(`[%+v on %s %d]`, message, file, line)
	}

	message = fmt.Sprintf(`[%+v]`, message)

	if logType == fatal {
		nativeLog.Fatalln(toString(``, color, message, params...))
	}

	if logType == err {
		nativeLog.Println(toString(``, color, message, params...))
		return
	}

	nativeLog.Println(toString(``, color, message, params...))
}
