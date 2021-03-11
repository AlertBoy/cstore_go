package main

import (
	"cstore/config"
	"cstore/gin"
)

func main() {
	config.Init()
	engine := server.NewRouter()

	engine.Run(":3030")
}
