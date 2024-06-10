package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func GetDb() *sql.DB {
	return db
}

func Init() {

	var err error
	db, err = sql.Open("sqlite", "local.db")
	if err != nil {

		fmt.Println(err)
		fmt.Println("===============")
		panic("No database connection")
	}
	db.SetMaxOpenConns(10)
	createTable()
}

func createTable() {
	createTableBook := ` 
	 CREATE TABLE IF NOT EXISTS books(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 title TEXT,
	 isbn TEXT,
	 author TEXT,
	 year INTEGER
	)`

	_, err := db.Exec(createTableBook)
	if err != nil {
		fmt.Println(err)
		fmt.Println("==================")
		panic("Cannot create table books")
	}
}
