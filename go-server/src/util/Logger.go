package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

var filename = "info.log"
var logger = &log.Logger{}

type LoggerHelp struct {
}

func initLog() {
	out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0)
	fmt.Println(filename)
	if err != nil {

	}
	logger = log.New(out, time.Now().String(), log.Ldate|log.Ltime|log.Llongfile)
}

func (loggerHelp *LoggerHelp) Info(message string) {
	fmt.Println("get log info ")
	initLog()
	logger.Println(message)
}

func (loggerHelp *LoggerHelp) Error(message string) {
	logger.Println(message)
}

func (loggerHelp *LoggerHelp) Debug(message string) {
	logger.Println(message)
}
func (loggerHelp *LoggerHelp) Warn(message string) {
	logger.Println(message)
}
func (loggerHelp *LoggerHelp) Fatal(message string) {
	logger.Println(message)
}
