package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Movies struct {
	id    int
	title string
	year  int
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	//result, err := db.Exec("insert into movies (title, year) VALUES ($1, $2)", "Lord of the rings 2", 2003)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//affected, _ := result.RowsAffected()
	//log.Printf("Rows affected %d", affected)

	// query for rows
	rows, err := db.Query("select * from movies")
	for rows.Next() {
		movies := Movies{}
		if err = rows.Scan(&movies.id, &movies.title, &movies.year); err != nil {
			log.Fatal(err)
		}
		log.Printf("movie is %v", movies)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// query for a single row
	var title string
	row := db.QueryRow("select title from movies where id=1")
	row.Scan(&title)
	log.Printf("movie with id 1: %s", title)
}
