package main

import (
	"log"
	"os"
	"sync"
)

type Logman struct {
	filename string
	*log.Logger
}

var theLog *Logman
var once sync.Once

// GetLogmanInstance : Get (or create) the instance of the logger.
func GetLogmanInstance() *Logman {
	once.Do(func() {
		theLog = CreateLogmanLogger("honeyshell.log")
	})

	return theLog
}

// CreateLogmanLogger : Create a logger.
func CreateLogmanLogger(fname string) *Logman {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	return &Logman{
		filename: fname,
		Logger:   log.New(file, "", log.Ldate|log.Ltime),
	}
}