package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/apis"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/middlewares"
)

func main() {
	// -------------- Init Echo -------------- //
	server := echo.New()
	// ---------------------------------------- //

	// -------------- Init Config -------------- //
	config.LoadEnv()
	// ---------------------------------------- //

	// -------------- Init Middlewares -------------- //
	middlewares.Init(server)
	// ---------------------------------------- //

	// -------------- Init APIs -------------- //
	apis.Init(server)
	// ---------------------------------------- //

	// -------------- Start Server -------------- //
	server.Logger.Fatal(server.Start(":8000"))
	// ---------------------------------------- //
}
