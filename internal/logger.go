package internal

import (
	"log"
	"net/http"
	"os"
	"sync"
	"fmt"
)

// TODO: Documenting, Logger 구현

type LogLevel int

type Logger struct {
	*log.Logger
	logLevel LogLevel
}

const (
	_ LogLevel = iota - 1
	DEBUG
	INFO
	WARN
	ERROR
	OFF
)

var logLevels = [...]string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"OFF",
}

var mu sync.RWMutex

func (l LogLevel) String() string {
	return logLevels[l]
}

func (l *Logger) SetLogLevel(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()
	l.logLevel = level
}

// Retrieve the current logging Level.
func (l *Logger) LogLevel() LogLevel {
	mu.RLock()
	defer mu.RUnlock()
	return l.logLevel
}

// Create a new *log.Logger wrapping w in a wlog.Writer
func NewLogger(fileName string, prefix string, flag int) *log.Logger {
	fpLog, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		return log.New(fpLog, prefix, flag)
	}
	return nil
}

func (l *Logger) Printf(level LogLevel, format string, v ...interface{}) {
	if level >= l.logLevel {
		formatted := level.String() + format
		l.Output(2, fmt.Sprintf(formatted, v...))
	}
}

func (l *Logger) Debug(format string, v ...interface{})  {
	l.Printf(DEBUG, format, v)
}

func (l *Logger) Info(format string, v ...interface{})  {
	l.Printf(INFO, format, v)
}

func (l *Logger) Warn(format string, v ...interface{})  {
	l.Printf(WARN, format, v)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.Printf(ERROR, format, v)
}

func WithLogger(l *Logger) Decorator {
	return func(c httpClient) httpClient {
		return handlerFunc(func(r *http.Request) (*http.Response, error) {
			l.Info(" %s", r.Method, r.URL)
			return c.Do(r)
		})
	}
}