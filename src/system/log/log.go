package log

import (
	"os"
	"reflect"

	"github.com/KayacChang/API_Server/system/env"
	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
)

// Fields type alias for log.Fields
type Fields = log.Fields

// Entry type alias for log.Entry
type Entry = log.Entry

func init() {

	if !env.IsDebug() {

		log.SetFormatter(&log.JSONFormatter{})
	}

	log.SetOutput(os.Stdout)
}

// WithFields Adds a struct of fields to the log entry.
func WithFields(arg interface{}) *Entry {

	var field Fields

	switch v := reflect.ValueOf(arg); v.Kind() {

	case reflect.Struct:
		field = structs.Map(arg)

	case reflect.Map:
		field = arg.(log.Fields)

	default:
		Fatal("unhandled kind %s", v.Kind())
	}

	return log.WithFields(field)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger
// 	then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the standard logger
// 	then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
