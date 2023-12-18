package utilities

import (
	"log/slog"
	"os"

	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func createLogFile(logPath string) *os.File {
	// Creating a directory if not exist
	err := os.MkdirAll("./logs", os.ModePerm)
	handler.ErrorHandler(err)

	// Creating a log file if not exist
	_, er := os.Create(logPath)
	handler.ErrorHandler(er)

	// Opening a log file so slog can write to it
	f, e := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	handler.ErrorHandler(e)
	return f
}

func InitApiLogs() *slog.Logger {
	pwd, _ := os.Getwd()
	logPath := pwd + config.K.String("api_logs")

	// Create the log file if it doesn't exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		f := createLogFile(logPath)
		defer f.Close()
		return slog.New(slog.NewJSONHandler(f, nil))
	}
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	handler.ErrorHandler(err)
	return slog.New(slog.NewJSONHandler(f, nil))
}

func InitErrLogs() *os.File {
	pwd, _ := os.Getwd()
	logPath := pwd + config.K.String("err_logs")

	// Create the log file if it doesn't exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		f := createLogFile(logPath)
		defer f.Close()
		return f
	}
	return nil
}
