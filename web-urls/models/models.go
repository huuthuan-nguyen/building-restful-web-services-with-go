package models

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgres://postgres:12345678@localhost/web_urls?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		// create model for our URL service
		stmt, err := db.Prepare("CREATE TABLE web_url(id SERIAL PRIMARY KEY, url TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}