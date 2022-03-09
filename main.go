package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var title string
	rows, err := db.Query("select title from movies")
	for rows.Next() {
		if err = rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		log.Printf("title is %s", title)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
