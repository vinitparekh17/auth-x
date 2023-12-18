package middlewares

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/utilities"
)

func Init(e *echo.Echo) {
	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `method=${method}, uri=${uri}, status=${status}, time=${latency_human}`,
		Output:           os.Stdout,
		CustomTimeFormat: "2006-01-02 05:04:05",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.K.Strings("origin"),
	}))

	e.Use(middleware.Secure())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
		Level: 5,
	}))

	e.Use(middleware.RequestID())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",`,
		Output: utilities.InitErrLogs(),
	}))

	e.Use(middleware.Recover())

	logger := utilities.InitApiLogs()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
}
