package utils

import (
	"io/ioutil"
	"log"
	"os"
)

type LoggingStruct struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

func NewLogging() *LoggingStruct {

	logging := &LoggingStruct{
		Trace: log.New(ioutil.Discard, "TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Info: log.New(os.Stdout, "INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(os.Stdout, "WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(os.Stderr, "ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Fatal: log.New(os.Stderr, "FATAL: ",
			log.Ldate|log.Ltime|log.Lshortfile),
	}

	return logging

}
