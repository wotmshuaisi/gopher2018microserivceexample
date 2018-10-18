package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger ...
func InitLogger() {
	var logFilePath = GetSysEnv("EXAMPLE_MAIN_LOG_PATH", "log/")

	logFile, err := os.OpenFile(logFilePath+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
		// logrus.SetOutput(os.Stdout)
	} else {
		logrus.Fatalln("Failed to write log to file", err)
	}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
	})
	return
}

// NewWebLogger ...
func NewWebLogger() *logrus.Logger {
	var logFilePath = GetSysEnv("EXAMPLE_WEB_LOG_PATH", "log/")

	logger := logrus.New()

	logFile, err := os.OpenFile(logFilePath+"web.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = io.MultiWriter(os.Stdout, logFile)
		// logrus.SetOutput(os.Stdout)
	} else {
		logrus.Fatalln("Failed to write weblog to file", err)
	}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logger.Formatter = &logrus.TextFormatter{
		DisableColors: false,
	}
	return logger
}
