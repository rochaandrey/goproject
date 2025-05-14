package main

import (
	"github.com/rochaandrey/goproject.git/config"
	"github.com/rochaandrey/goproject.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
