/*
Copyright © 2021 GUSTAVO SILVA <gustavosantaremsilva@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// InitLogger instantiates three types of logs: An info log, a warn log and an error log.
func InitLogger() *Log {
	return &Log{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile),
	}
}

func (l *Log) Info(msg string) {
	l.InfoLogger.Print(msg)
}
func (l *Log) Error(msg string) {
	l.ErrorLogger.Fatalln(msg)
}
