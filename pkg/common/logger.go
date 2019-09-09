package common

import (
	"fmt"
	"log"
	"os"

	e "../e"
)

type Logger struct {
	logFileName string
	logFilePath string
	logLevel    int
}

var logger *Logger

func GetLogger() *Logger {
	if logger != nil {
		return logger
	}
	return newLogger()
}

func newLogger() *Logger {
	logger = new(Logger)
	logger.setLogFileName(GetCurrentTimeInRFC())
	logger.logLevel = SHOW_ERROR_AND_TRACK
	return logger
}

func (logger *Logger) SetLogFilePath(logFilePath string) {
	logger.logFilePath = logFilePath

}

func (logger *Logger) setLogFileName(fileName string) {
	logger.logFileName = fileName + LOG_FILE_FORMAT
}

func (logger Logger) stdPrint(level int, a ...interface{}) {
	switch logger.logLevel {
	case SHOW_ERROR_AND_TRACK:
		if level == e.TRACK {
			fmt.Print(DEBUG_MSG_TRACK)
		} else if level == e.ERROR {
			fmt.Print(DEBUG_MSG_ERROR)
		}
		fmt.Print(" ")
		for _, x := range a {
			fmt.Print(x)
		}
		fmt.Println()
	case SHOW_TRACK:
		if level == e.TRACK {
			fmt.Print(DEBUG_MSG_TRACK)
			fmt.Print(" ")
			for _, x := range a {
				fmt.Print(x)
			}
			fmt.Println()
		}
	}
}

func (logger Logger) Log(level int, a ...interface{}) {
	logger.stdPrint(level, a)
	logger.writeInLogFile(level, a)
}

func (logger Logger) writeInLogFile(level int, a ...interface{}) {
	f, err := os.OpenFile(logger.logFilePath+logger.logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Terminate(err)
	}
	defer f.Close()
	log.SetOutput(f)

	switch logger.logLevel {
	case SHOW_ERROR_AND_TRACK:
		if level == e.TRACK {
			log.Print(DEBUG_MSG_TRACK, ":", a)
		} else if level == e.ERROR {
			log.Print(DEBUG_MSG_ERROR, ":", a)
		}
	case SHOW_TRACK:
		if level == e.TRACK {
			log.Print(DEBUG_MSG_TRACK, ":", a)
		}
	}
}
