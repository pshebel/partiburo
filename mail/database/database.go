package database

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/pshebel/partiburo/mail/env"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", env.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// run migrations
	log.Printf("database initialized")
}

func GetDB() (*sql.DB, error) {
	fmt.Println(env.DB)
	return sql.Open("sqlite3", env.DB)
}