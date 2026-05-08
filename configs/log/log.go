package log

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	prefix  string
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

type Option func(*Logger)

func WithPrefix(p string) Option {
	return func(l *Logger) {
		if p != "" {
			l.prefix = "[" + p + "]"
		} else {
			l.prefix = ""
		}
	}
}

func NewLogger(opts ...Option) *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime

	l := &Logger{
		writer: writer,
		prefix: "",
	}

	for _, opt := range opts {
		opt(l)
	}

	prefix := l.prefix
	if prefix != "" {
		prefix += " "
	}

	l.debug = log.New(writer, prefix+"DEBUG: ", flags)
	l.info = log.New(writer, prefix+"INFO: ", flags)
	l.warning = log.New(writer, prefix+"WARNING: ", flags)
	l.err = log.New(writer, prefix+"ERROR: ", flags)
	l.fatal = log.New(writer, prefix+"FATAL: ", flags)

	return l
}

func (l *Logger) Debug(format string, v ...any) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Info(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warn(format string, v ...any) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...any) {
	l.err.Fatalf(format, v...)
}
