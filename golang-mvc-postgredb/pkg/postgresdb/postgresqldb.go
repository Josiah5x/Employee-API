package postgresdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func New(url string) *sql.DB {
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected by Josiah")
	return db

}
