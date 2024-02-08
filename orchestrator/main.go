package main

import (
	"fmt"
	"orchestrator/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Print(cfg)

	//todo: init slog
	//todo: init storage
	//todo: init router
	//todo: run server
}
