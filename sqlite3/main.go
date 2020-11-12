package main

import (
	"log"
	"database/sql"
	"os"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	id int
	name string
	author string
}

var db *sql.DB

func main() {
	var err error
	os.Remove("./books.db")
	db, err = sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// create table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)"); err != nil {
		log.Println("Error in creating table")
		fmt.Printf("%v\n", err)
	} else {
		log.Println("Successfully created table books!")
	}
	
	// insert
	statement, _ := db.Prepare("INSERT INTO books(name, author, isbn) VALUES (?, ?, ?)")
	statement.Exec("A tale of 2 cites", "Charles Dickens", 140430547)
	log.Println("Successfully inserted the book into database!")
	// read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// update
	statement, _ = db.Prepare("UPDATE books SET name=? WHERE id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")
	// delete
	statement, _ = db.Prepare("DELETE FROM books WHERE id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}