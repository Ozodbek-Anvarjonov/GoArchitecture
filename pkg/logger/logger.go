package logger

import (
	"log"
	"os"
	"strings"
)

type Logger struct {
	level string
}

func New(level string) *Logger {
	return &Logger{
		level: strings.ToLower(level),
	}
}

func (l *Logger) Info(msg string) {
	if l.level == "info" || l.level == "debug" {
		log.Println("[INFO] " + msg)
	}
}

func (l *Logger) Debug(msg string) {
	if l.level == "debug" {
		log.Println("[DEBUG] " + msg)
	}
}

func (l *Logger) Error(err error) {
	log.Println("[ERROR]", err)
}

func (l *Logger) Fatal(err error) {
	log.Println("[FATAL]", err)
	os.Exit(1)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *Logger) Println(v ...interface{}) {
	log.Println(v...)
}
