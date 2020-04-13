package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const logDir = "flight-logs/"

func init() {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModeDir|0755)
	}

	file, err := os.OpenFile(logDir + "main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
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

// LogWarning prints and logs a string
func LogWarning(str string) {
	fmt.Println(str)
	logrus.Warning(str)
}

// LogInfo prints and logs a string
func LogInfo(str string) {
	fmt.Println(str)
	logrus.Info(str)
}
