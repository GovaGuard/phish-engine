package main

import (
	"flag"
	"log"
	"os"

	"github.com/holgerson97/phish-engine/internal/db"
	"github.com/holgerson97/phish-engine/internal/router"
	"github.com/holgerson97/phish-engine/internal/usecase"
	"github.com/holgerson97/phish-engine/pkg/info"
)

func main() {
	rethinkDB := flag.String("rethink-db", "localhost:28015", "The address of the RethinkDB server.")
	v := flag.Bool("version", false, "Prints the version of phish-engine.")
	flag.Parse()

	if *v {
		info.PrintVersion()
		os.Exit(0)
	}

	_, err := db.New(*rethinkDB)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB Connected")

	usecase, err := usecase.New(*rethinkDB)
	if err != nil {
		return
	}

	router := router.New(usecase)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

	log.Print("Server Running")
}
