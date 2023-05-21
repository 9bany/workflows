package main

import (
	"github.com/9bany/workflows/server"
	"github.com/9bany/workflows/utils"
)

func main() {
	config := utils.LodConfig(".")
	serve := server.NewServer(config)
	serve.Start()
}
