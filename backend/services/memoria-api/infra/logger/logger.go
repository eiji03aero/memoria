package logger

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	mu     sync.Mutex
	level  LogLevel
	logger *log.Logger
}

func NewLogger(level LogLevel, out io.Writer) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(out, "", log.Ldate|log.Ltime),
	}
}

func (l *Logger) logMessage(level LogLevel, levelStr string, messages ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level >= l.level {
		l.logger.SetPrefix(fmt.Sprintf("%s [%s] ", time.Now().Format("2006-01-02 15:04:05"), levelStr))
		l.logger.Println(messages...)
	}
}

func (l *Logger) Debug(messages ...any) {
	l.logMessage(DEBUG, "DEBUG", messages...)
}

func (l *Logger) Info(messages ...any) {
	l.logMessage(INFO, "INFO", messages...)
}

func (l *Logger) Warn(messages ...any) {
	l.logMessage(WARN, "WARN", messages...)
}

func (l *Logger) Error(messages ...any) {
	l.logMessage(ERROR, "ERROR", messages...)
}

func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) GetLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}
