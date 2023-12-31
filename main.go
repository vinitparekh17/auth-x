package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/apis"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/middlewares"
	srv "github.com/vinitparekh17/project-x/server"
)

type Server echo.Echo

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

	// ---------------------------------------- //
	srv.Init(server)
}
