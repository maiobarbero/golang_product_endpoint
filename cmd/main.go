package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := app{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Printf("server failed to start. %s", err)
		os.Exit(1)
	}
}
