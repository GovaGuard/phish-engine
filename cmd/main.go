package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/holgerson97/phish-engine/internal/mail"
	"github.com/holgerson97/phish-engine/internal/router"
	"github.com/holgerson97/phish-engine/internal/usecase"
	"github.com/holgerson97/phish-engine/pkg/info"
	"github.com/holgerson97/phish-engine/repository/rethinkdb"
)

func main() {
	rethinkDB := flag.String("rethink-db", "localhost:28015", "The address of the RethinkDB server.")
	port := flag.Int("port", 8080, "")

	sender := flag.String("smtp-sender", "", "")
	username := flag.String("smtp-username", "", "")
	password := flag.String("smtp-password", "", "")
	smtpServer := flag.String("smtp-smtp-server", "", "")
	smtpAuth := flag.String("smtp-smtp-auth", "", "")

	v := flag.Bool("version", false, "Prints the version of phish-engine.")
	flag.Parse()

	if *v {
		info.PrintVersion()
		os.Exit(0)
	}

	log.Print("DB Connected")

	// TODO: call constructor
	// Issue URL: https://github.com/GovaGuard/phish-engine/issues/11
	m := mail.Sender{
		Sender:     *sender,
		User:       *username,
		Password:   *password,
		SMTPServer: *smtpServer,
		Host:       *smtpAuth,
	}

	repo, err := rethinkdb.NewClient(context.TODO(), *rethinkDB)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	usecase := usecase.New(repo, repo, m)
	if err != nil {
		return
	}

	router := router.New(usecase)
	if err := router.Run(*port); err != nil {
		log.Fatal(err)
	}

	log.Print("Server Running")
}
