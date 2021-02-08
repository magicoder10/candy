package observability

import (
    "fmt"
    "runtime"
    "time"
)

type LogLevel int

const (
    Off   LogLevel = 0
    Fatal LogLevel = 1
    Error LogLevel = 2
    Warn  LogLevel = 3
    Info  LogLevel = 4
    Debug LogLevel = 5
    Trace LogLevel = 6
)

var logLevelNames = map[LogLevel]string{
    Fatal: "FATAL",
    Error: "ERROR",
    Warn:  "WARN",
    Info:  "INFO",
    Debug: "DEBUG",
    Trace: "TRACE",
}

type Logger struct {
    logLevel LogLevel
}

// Errors which causes service or application shutdown, usually with data corruption.
func (l Logger) Fatalf(format string, arr ...interface{}) {
    l.logWithLineTrace(Fatal, format, ANSIRed, arr...)
}

func (l Logger) Fatatln(message string) {
    l.Fatalf("%s\n", message)
}

// Fatal errors to certain operations but now the whole service or app.
func (l Logger) Errorf(err error) {
    l.logWithLineTrace(Error, "%w", ANSIRed, err)
}

func (l Logger) Errorln(err error) {
    l.Errorf(err)
}

// Potential issues with automatic recovery.
func (l Logger) Warnf(format string, arr ...interface{}) {
    l.logf(Warn, format, ANSIYellow, arr...)
}

func (l Logger) Warnln(message string) {
    l.Warnf("%s\n", message)
}

// Generally useful information.
func (l Logger) Infof(format string, arr ...interface{}) {
    l.logf(Info, format, ANSIDefault, arr...)
}

func (l Logger) Infoln(message string) {
    l.Infof("%s\n", message)
}

// Diagnostically helpful to people more than just developers.
func (l Logger) Debugf(format string, arr ...interface{}) {
    l.logf(Debug, format, ANSIDefault, arr...)
}

func (l Logger) Debugln(message string) {
    l.Debugf("%s\n", message)
}

// Trace the part of the code in function level.
func (l Logger) Tracef(format string, arr ...interface{}) {
    l.logWithLineTrace(Trace, format, ANSIDefault, arr...)
}

func (l Logger) Traceln(message string) {
    l.Tracef("%s\n", message)
}

func (l Logger) isVisible(logLevel LogLevel) bool {
    return logLevel <= l.logLevel
}

func (l Logger) logWithLineTrace(level LogLevel, format string, color string, arr ...interface{}) {
    newFormat := fmt.Sprintf("%s\n%s", lineTrace(), format)
    l.logf(level, newFormat, color, arr...)
}

func withColor(format string, color string) string {
    return fmt.Sprintf("%s%s%s", color, format, ANSIReset)
}

func (l Logger) logf(level LogLevel, format string, color string, arr ...interface{}) {
    if !l.isVisible(level) {
        return
    }

    now := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf(withColor("%s [%s] %s", color), now, logLevelNames[level], fmt.Sprintf(format, arr...))
}

func lineTrace() string {
    _, file, line, _ := runtime.Caller(3)
    return fmt.Sprintf("%s:%d", file, line)
}

func NewLogger(logLevel LogLevel) Logger {
    return Logger{logLevel: logLevel}
}
