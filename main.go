package main

import (
	"distivity/config"
	"distivity/server"
)

func main() {
	config := config.GetConfig()
	server.Run(config)
}
