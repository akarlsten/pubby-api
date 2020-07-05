package main

import (
	"github.com/akarlsten/pubby-api/app"
	"github.com/akarlsten/pubby-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
