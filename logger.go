package logger

import (
	"encoding/json"
	"fmt"
	"io"
)

type Logger struct {
	w io.Writer
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{w: w}
}

type Entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

func (l *Logger) Print(m string, s string) {
	entry := Entry{
		Message:  m,
		Severity: s,
	}
	_ = json.NewEncoder(l.w).Encode(entry)
}

func (l *Logger) Debug(m string) {
	l.Print(m, "debug")
}

func (l *Logger) Debugf(m string, a ...interface{}) {
	l.Debug(fmt.Sprintf(m, a...))
}

func (l *Logger) Info(m string) {
	l.Print(m, "info")
}

func (l *Logger) Infof(m string, a ...interface{}) {
	l.Info(fmt.Sprintf(m, a...))
}

func (l *Logger) Notice(m string) {
	l.Print(m, "notice")
}

func (l *Logger) Noticef(m string, a ...interface{}) {
	l.Notice(fmt.Sprintf(m, a...))
}

func (l *Logger) Warn(m string) {
	l.Print(m, "warning")
}

func (l *Logger) Warnf(m string, a ...interface{}) {
	l.Warn(fmt.Sprintf(m, a...))
}

func (l *Logger) Error(m string) {
	l.Print(m, "error")
}

func (l *Logger) Errorf(m string, a ...interface{}) {
	l.Error(fmt.Sprintf(m, a...))
}

func (l *Logger) Critical(m string) {
	l.Print(m, "critical")
}

func (l *Logger) Criticalf(m string, a ...interface{}) {
	l.Critical(fmt.Sprintf(m, a...))
}

func (l *Logger) Alert(m string) {
	l.Print(m, "alert")
}

func (l *Logger) Alertf(m string, a ...interface{}) {
	l.Alert(fmt.Sprintf(m, a...))
}
