package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Movies struct {
	ID        int
	Title     string
	Year      int
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	db, err := sqlx.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	movies := []Movies{}
	err = db.Select(&movies, "select * from movies")
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		log.Printf("Movie is %v", movie.Title)
	}
}
