package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	file, err := os.OpenFile("log/", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}

// LogFatalAndExitIfNotNull prints an logs and error and exits if err not null
func LogFatalAndExitIfNotNull(err error) {
	if err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
		os.Exit(1)
	}
}

// LogErrorIfNotNull prints and logs an error if err not null
func LogErrorIfNotNull(err error) {
	if err != nil {
		fmt.Println(err)
		logrus.Error(err)
	}
}

// LogInfo prints and logs a string
func LogInfo(str string) {
	fmt.Println(str)
	logrus.Info(str)
}
