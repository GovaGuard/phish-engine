package main

import (
	"log"

	"github.com/holgerson97/phish-engine/internal/router"
)

func main() {
	log.Print("Server Running")
	if err := router.Run(); err != nil {
		panic(err)
	}
}
