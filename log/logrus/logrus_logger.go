package logrus

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	//Log is
	Log *logrus.Logger
)

const (
	//Loglevel is
	Loglevel = "info"
)

func init() {
	level, err := logrus.ParseLevel(Loglevel)

	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}

//Info is
func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}

//Error is
func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("Error")
	Log.WithFields(parseFields(tags...)).Error(msg)
}

//Debug is
func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}
