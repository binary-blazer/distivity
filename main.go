package main

import (
	"distivity/config"
	"distivity/routes"
	"distivity/server"
)

func main() {
	config := config.GetConfig(routes.IndexHandler)
	server.Run(config, routes.IndexHandler)
}
