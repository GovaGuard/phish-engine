package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/holgerson97/phish-engine/internal/mail"
	"github.com/holgerson97/phish-engine/internal/router"
	"github.com/holgerson97/phish-engine/internal/usecase"
	"github.com/holgerson97/phish-engine/pkg/info"
	"github.com/holgerson97/phish-engine/repository/mongodb"
	"golang.org/x/sync/errgroup"
)

func main() {
	db := flag.String("db", "mongodb://phish:phish@localhost:27017", "The address of the DB server.")
	port := flag.Int("port", 8080, "")

	sender := flag.String("smtp-sender", "", "")
	username := flag.String("smtp-username", "", "")
	password := flag.String("smtp-password", "", "")
	smtpServer := flag.String("smtp-server", "", "")
	smtpAuth := flag.String("smtp-auth", "", "")

	v := flag.Bool("version", false, "Prints the version of phish-engine.")
	flag.Parse()

	if *v {
		info.PrintVersion()
		os.Exit(0)
	}

	// TODO: call constructor
	// TODO: Add HealthCheck for SMTP Server
	m := mail.Sender{
		Sender:     *sender,
		User:       *username,
		Password:   *password,
		SMTPServer: *smtpServer,
		Host:       *smtpAuth,
	}

	repo, err := mongodb.NewClient(context.TODO(), *db)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	log.Print("DB Connected")

	usecase := usecase.New(repo, repo, m)
	if err != nil {
		return
	}

	group := new(errgroup.Group)
	group.Go(func() error {
		for {
			log.Print("iterating campaigns")
			if err := usecase.WorkCampaigns(); err != nil {
				log.Print(err)
			}

			time.Sleep(10 * time.Second)
		}
	})

	router := router.New(usecase)
	if err := router.Run(*port); err != nil {
		log.Fatal(err)
	}

	log.Print("Server Running")
}
