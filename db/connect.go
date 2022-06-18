package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection_url := "user=postgres dbname=lojinha password=postgres host=localhost sslmode=disable"
	conn, err := sql.Open("postgres", connection_url)

	if err != nil {
		panic(err.Error())
	}
	return conn
}
