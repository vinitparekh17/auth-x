package handler

import "log/slog"

func ErrorHandler(e error) {
	if e != nil {
		slog.Error(e.Error())
	}
}
