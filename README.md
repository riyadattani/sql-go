# Postgres SQL and Go

1. Start up Postgres `docker-compose --build up postgres`
2. Create a table in the DB
3. Run `go run main.go`

Tutorial from O'Reilly: SQL and Go by Mark Bates 

Packages:
- github.com/jmoiron/sqlx - this is package that sits on top of `database/sql`
- github.com/lib/pq