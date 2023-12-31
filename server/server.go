package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func Init(e *echo.Echo) {
	var srv http.Server
	// -------------- Get Port -------------- //
	p, er := config.GetEnv("PORT")
	handler.ErrorHandler(er)
	// ---------------------------------------- //

	// -------------- Init Server -------------- //
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt) // Notify the channel if an interrupt signal is received

	// This go routine will start the server in a separate thread so that we can listen to the interrupt signal in the main thread
	go func() {
		e.Server = &srv
		e.Start(":" + p)
	}()

	<-sigint      // this line will block the main thread until we receive a signal
	close(sigint) // closing the channel so that we can exit gracefully

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// -------------- Gracefull Shutdown -------------- //
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		slog.Error("Gracefull Shutdown failed: %v", err)
		os.Exit(1)
	}
	slog.Info("\nGracefull Shutdown successfull")
	// ---------------------------------------- //
}
