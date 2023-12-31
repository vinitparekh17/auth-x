package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/apis"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/middlewares"
	"github.com/vinitparekh17/project-x/redis"
	srv "github.com/vinitparekh17/project-x/server"
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

	// -------------- Init Redis -------------- //
	redis.Init()
	// ---------------------------------------- //

	// -------------- Init Server -------------- //
	srv.Init(server)
	// ---------------------------------------- //
}
