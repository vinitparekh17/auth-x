package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/apis"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/handler"
	"github.com/vinitparekh17/project-x/middlewares"
)

func main() {
	// -------------- Init Echo -------------- //
	server := echo.New()
	// ---------------------------------------- //

	// -------------- Load Env -------------- //
	config.LoadEnv()
	// ---------------------------------------- //

	// -------------- Load Config -------------- //
	config.LoadConfig()
	// ---------------------------------------- //

	// -------------- Init Database -------------- //
	database.Init()
	// ---------------------------------------- //

	// -------------- Init Middlewares -------------- //
	middlewares.Init(server)
	// ---------------------------------------- //

	// -------------- Init APIs -------------- //
	apis.Init(server)
	// ---------------------------------------- //

	// -------------- Get Port -------------- //
	p, e := config.GetConfig("PORT")
	handler.ErrorHandler(e)
	// ---------------------------------------- //

	// -------------- Start Server -------------- //
	server.Logger.Fatal(server.Start(":" + p))
	// ---------------------------------------- //
}
