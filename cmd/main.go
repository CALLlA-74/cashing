package main

import (
	"github.com/CALLlA-74/cashing/internal/app"
	"github.com/CALLlA-74/cashing/internal/config"
)

func main() {
	cfg := config.GetConfig()
	app.New(cfg).Start()
}
