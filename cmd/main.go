package main

import (
	"log"

	"github.com/holgerson97/phish-engine/internal/db"
	"github.com/holgerson97/phish-engine/internal/router"
)

func main() {
	log.Print("DB Connected")
	_, err := db.New("localhost:28015")
	if err != nil {
		panic(err)
	}

	log.Print("Server Running")
	if err := router.Run(); err != nil {
		panic(err)
	}
}
