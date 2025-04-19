package main

import (
	"log"

	"github.com/holgerson97/phish-engine/internal/db"
	"github.com/holgerson97/phish-engine/internal/router"
	"github.com/holgerson97/phish-engine/internal/usecase/campaigns"
)

func main() {
	log.Print("DB Connected")
	_, err := db.New("localhost:28015")
	if err != nil {
		panic(err)
	}

	usecase, err := campaigns.New("localhost:28015")
	if err != nil {
		return
	}

	router := router.New(usecase)
	if err := router.Run(); err != nil {
		panic(err)
	}

	log.Print("Server Running")
}
