package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type LoggerLevel interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	Panic(message string)
}

type logger struct {
	log *log.Entry
}

func NewLog(eventName string) LoggerLevel {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)

	l := log.WithFields(log.Fields{
		"event_name": eventName,
	})

	return &logger{
		log: l,
	}
}

func (l *logger) Debug(message string) {
	l.log.WithFields(log.Fields{"message": message}).Debug()
}

func (l *logger) Info(message string) {
	l.log.WithFields(log.Fields{"message": message}).Info()
}

func (l *logger) Warn(message string) {
	l.log.WithFields(log.Fields{"message": message}).Warn()
}

func (l *logger) Error(message string) {
	l.log.WithFields(log.Fields{"message": message}).Error()
}

func (l *logger) Fatal(message string) {
	l.log.WithFields(log.Fields{"message": message}).Fatal()
}

func (l *logger) Panic(message string) {
	l.log.WithFields(log.Fields{"message": message}).Panic()
}
