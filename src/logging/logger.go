package logging

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Level defines all available log levels for log messages.
type Level int

// Log levels.
const (
	ERROR Level = iota
	WARNING
	INFO
	DEBUG
)

var levelNames = []string{
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

//Logger
var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

//actLevel returns the actual Log Level
var actLevel Level

func GetActLevel() Level {
	return actLevel
}

// LogLevel returns the log level from a string representation.
func logLevel(level string) (Level, error) {
	for i, name := range levelNames {
		if strings.EqualFold(name, level) {
			return Level(i), nil
		}
	}
	return ERROR, errors.New("logger: invalid log level")
}

//Creates individual loggers for each log level
//The given handle defines the output channel of the logger
// - ioutil.Discard for switch off the logger
// - os.Stdout or os.Stderr for writing to console
func create(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	Debug = log.New(traceHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Init initialize the logger depends on the given log level.
//Log Level "ERROR", "WARNING",	"INFO",	"DEBUG"
func Init(loglevel string) error {
	var err error
	actLevel, err = logLevel(loglevel)
	if err != nil {
		return err
	}

	if actLevel == ERROR {
		create(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	}
	if actLevel == WARNING {
		create(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
	}
	if actLevel == INFO {
		create(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
	if actLevel == DEBUG {
		create(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	}

	return err
}
