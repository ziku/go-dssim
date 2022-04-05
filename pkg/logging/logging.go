package logging

import "log"

var (
	LogFatal *log.Logger
	LogError *log.Logger
	LogWarn  *log.Logger
	LogInfo  *log.Logger
	LogDebug *log.Logger
)

func init() {
	LogFatal = log.New(log.Writer(), "FATAL: ", log.Ldate|log.Ltime)
	LogError = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime)
	LogWarn = log.New(log.Writer(), "WARN: ", log.Ldate|log.Ltime)
	LogInfo = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime)
	LogDebug = log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime)
}

func Fatal(str string, v ...interface{}) {
	LogFatal.Printf(str, v...)
}

func Error(str string, v ...interface{}) {
	LogError.Printf(str, v...)
}

func Warn(str string, v ...interface{}) {
	LogWarn.Println(v...)
}

func Info(str string, v ...interface{}) {
	LogInfo.Printf(str, v...)
}

func Debug(str string, v ...interface{}) {
	LogDebug.Println(v...)
}
