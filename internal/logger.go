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

type Log struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// InitLogger instantiates a struct that provides logging abilities with predefined formats.
func InitLogger() *Log {
	return &Log{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info does a basic STDOUT logging. No side effects on this call.
func (l *Log) Info(msg string) {
	l.InfoLogger.Println(msg)
}

// Error calls Fatal, meaning it logs and exists after call. Use it carefully.
func (l *Log) Error(msg string) {
	l.ErrorLogger.Fatalln(msg)
}
