package logs

import (
	"log"
	"os"
)

// Log is the struct that contains all types of logs used in the application. The motivation for this custom approach
// is to standardise all logs in the application and also facilitate introducing different verbosity levels.
// InfoLogger can be used to replace all fmt.Println calls.
// ErrorLogger can be used to call fatal after logging.
type Log struct {
	infoLogger *log.Logger
	errorLog   *log.Logger
}

// InitLogger instantiates three types of logs: An info log, a warn log and an error log.
func InitLogger() *Log {
	return &Log{
		infoLogger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		errorLog:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile),
	}
}

func (l *Log) Info(msg string) {
	l.infoLogger.Print(msg)
}
func (l *Log) Error(msg string) {
	l.errorLog.Fatalln(msg)
}
