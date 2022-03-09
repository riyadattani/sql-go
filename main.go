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

	//transaction
	tx := db.MustBegin()
	m := Movies{
		Title: "Batman",
		Year:  2022,
	}
	tx.Exec("insert into movies (title, year) values ($1, $2)", m.Title, m.Year)
	tx.Commit()

	movies := []Movies{}
	err = db.Select(&movies, "select * from movies")
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		log.Printf("Movie is %v", movie.Title)
	}
}
