package main

import (
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/server"
)

func main() {
	config.Init()
	server.Init()
}
