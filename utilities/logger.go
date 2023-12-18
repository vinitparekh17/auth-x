package utilities

import (
	"log/slog"
	"os"

	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func InitApiLogs() *slog.Logger {
	pwd, _ := os.Getwd()
	logPath := pwd + config.K.String("api_log")

	// Create the log file if it doesn't exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		// Creating a directory if not exist
		err := os.MkdirAll("./logs", os.ModePerm)
		handler.ErrorHandler(err)

		// Creating a log file if not exist
		_, er := os.Create(logPath)
		handler.ErrorHandler(er)

		// Opening a log file so slog can write to it
		f, e := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		handler.ErrorHandler(e)
		defer f.Close()
		return slog.New(slog.NewJSONHandler(f, nil))
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func InitErrLogs() *os.File {
	pwd, _ := os.Getwd()
	logPath := pwd + config.K.String("err_log")

	// Create the log file if it doesn't exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		// Creating a directory if not exist
		err := os.MkdirAll("./logs", os.ModePerm)
		handler.ErrorHandler(err)

		// Creating a log file if not exist
		_, er := os.Create(logPath)
		handler.ErrorHandler(er)

		// Opening a log file so slog can write to it
		f, e := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		handler.ErrorHandler(e)
		defer f.Close()
		return f
	}
	return nil
}
