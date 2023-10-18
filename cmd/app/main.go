package main

import (
	"log"

	"todo2"
	"todo2/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo2.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
