package util

import (
	"log"
	"os"
	"time"
)

var filename = "info.log"
var logger = &log.Logger{}
type LoggerHelp struct {
}
func init(){
	out,err := os.Open(filename)
	if err != nil {

	}
	logger = log.New(out,time.Now().String(),log.Ldate|log.Ltime|log.Llongfile)
}

func (loggerHelp *LoggerHelp) Info(message string){
	 logger.Println(message)
}

func (loggerHelp *LoggerHelp) Error(message string){
	logger.Println(message)
}

func (loggerHelp *LoggerHelp) Debug(message string){
	logger.Println(message)
}
func (loggerHelp *LoggerHelp) Warn(message string){
	logger.Println(message)
}
func (loggerHelp *LoggerHelp) Fatal(message string){
	logger.Println(message)
}



