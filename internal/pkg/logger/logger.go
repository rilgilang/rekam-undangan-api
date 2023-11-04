package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

//log.Trace("Something very low level.")
//log.Debug("Useful debugging information.")
//log.Info("Something noteworthy happened!")
//log.Warn("You should probably take a look at this.")
//log.Error("Something failed but I'm not quitting.")
//// Calls os.Exit(1) after logging
//log.Fatal("Bye.")
//// Calls panic() after logging
//log.Panic("I'm bailing.")

type LoggerLevel interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	Panic(message string)
}

type logger struct {
	logs *log.Entry
}

func NewLog(eventName string) LoggerLevel {
	logs := log.WithFields(log.Fields{
		"event_name": eventName,
	})

	return &logger{
		logs: logs,
	}
}

func (l *logger) Debug(message string) {
	l.Debug(message)
}

func (l *logger) Info(message string) {
	l.Info(message)
}

func (l *logger) Warn(message string) {
	l.Warn(message)
}

func (l *logger) Error(message string) {
	l.Error(message)
}

func (l *logger) Fatal(message string) {
	l.Fatal(message)
}

func (l *logger) Panic(message string) {
	l.Panic(message)
}
