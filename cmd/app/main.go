package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/csv"

	"github.com/alexvitayu/gorm-project/internal/config"
	"github.com/alexvitayu/gorm-project/internal/db"
	"github.com/alexvitayu/gorm-project/internal/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: movies <list|create|show|update|delete> [args]")
	}

	cfg := config.Load()

	db, err := db.Open(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	entity := os.Args[1]
	action := os.Args[2]
	if entity != "movies" {
		log.Fatal("only movies supported")
	}

	switch action {
	case "list":
		handleList(db)
	case "create":
		handleCreate(db, os.Args)
	case "show":
		handleShow(db, os.Args)
	case "update":
		handleUpdate(db, os.Args)
	case "delete":
		handleDelete(db, os.Args)
	default:
		log.Fatal("unknown action")
	}
}

func handleList(db *gorm.DB) {
	var movies []models.Movie
	if err := db.Find(&movies).Error; err != nil {
		log.Printf("no movies found: %v", err)
	}
	log.Printf("movies: %d", len(movies))
}

func handleCreate(db *gorm.DB, args []string) {
	if len(args) < 7 {
		log.Fatal("usage: movies create <title> <genre> <released_at> <description> <rating>")
	}

	layout := "2006-01-02"
	t, err := time.Parse(layout, args[5])
	if err != nil {
		log.Fatalf("fail to convert to time: %v", err)
	}
	r, err := decimal.NewFromString(args[7])
	if err != nil {
		log.Fatalf("fail to convert to decimal: %v", err)
	}

	movie := models.Movie{
		Title:       args[3],
		Genre:       args[4],
		ReleasedAt:  t,
		Description: args[6],
		Rating:      r,
	}
	if err = db.Create(&movie).Error; err != nil {
		log.Fatalf("fail to creare movie: %v", err)
	}
	log.Printf("created movie with id = %d", movie.ID)
}

func handleShow(db *gorm.DB, args []string) {
	var movie models.Movie
	if err := db.First(&movie, args[3]).Error; err != nil {
		log.Fatalf("fail to show movie: %v", err)
	}
	log.Printf("movie_title=%v", movie.Title)

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	if err := writer.Write([]string{"id", "title", "genre", "released_at", "description", "rating"}); err != nil {
		log.Fatalf("fail to write header: %v", err)
	}

	record := []string{
		fmt.Sprint(movie.ID),
		movie.Title,
		movie.Genre,
		fmt.Sprint(movie.ReleasedAt),
		movie.Description,
		fmt.Sprint(movie.Rating),
	}

	if err := writer.Write(record); err != nil {
		log.Fatalf("fail to write body: %v", err)
	}
}

func handleUpdate(db *gorm.DB, args []string) {
	if len(args) < 6 {
		log.Fatal("usage: movies update <id> <field> <value>")
	}
	if err := db.Model(&models.Movie{}).
		Where("id = ?", args[3]).
		Update(args[4], args[5]).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("movie updated")
}

func handleDelete(db *gorm.DB, args []string) {
	res := db.Delete(&models.Movie{}, args[3])
	if res.Error != nil {
		log.Fatal("fail to delete movie")
	}
	log.Println("movie deleted")
	log.Printf("rows_affected = %v\n", res.RowsAffected)
}
