package main

import (
	"log"

	"github.com/personalhub/personalhub-back/src/config"
	"github.com/personalhub/personalhub-back/src/router"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	router.Run()
}
