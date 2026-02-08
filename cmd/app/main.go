package main

import (
	"log"
	"os"

	"github.com/alexvitayu/gorm-project/internal/app"
	"github.com/alexvitayu/gorm-project/internal/config"
	"github.com/alexvitayu/gorm-project/internal/db"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: movies <list|create|show|update|delete|unrated> [args]")
	}

	cfg := config.Load()

	conn, err := db.Open(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	entity := os.Args[1]
	action := os.Args[2]
	if entity != "movies" {
		log.Fatal("only movies supported")
	}

	switch action {
	case "list":
		app.HandleList(conn, os.Args)
	case "create":
		app.HandleCreate(conn, os.Args)
	case "show":
		app.HandleShow(conn, os.Args)
	case "update":
		app.HandleUpdate(conn, os.Args)
	case "delete":
		app.HandleDelete(conn, os.Args)
	case "unrated":
		app.HandleUnrated(conn)
	default:
		log.Fatal("unknown action")
	}
}
