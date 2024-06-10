package models

import (
	"BookWebApi/db"
	"database/sql"
	"fmt"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (b *Book) Save() error {
	fields := []string{"title", "isbn", "author", "year"}
	values := []interface{}{b.Title, b.Isbn, b.Author, b.Year}

	result, err := db.Insert("books", fields, values)
	if err != nil {
		return err
	}

	bookId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	b.Id = bookId
	return nil
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Select("books", Book{})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Title, &book.Isbn, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	row := db.GetDb().QueryRow("SELECT * FROM books WHERE id = ?", id)
	var book Book
	err := row.Scan(&book.Id, &book.Title, &book.Isbn, &book.Author, &book.Year)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Book not found")
	} else if err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(book Book) error {
	fields := []string{"title", "isbn", "author", "year"}
	values := []interface{}{book.Title, book.Isbn, book.Author, book.Year}
	_, err := db.Update("books", fields, values, book.Id)
	return err
}

func DeleteBook(id int64) error {
	_, err := db.Delete("books", id)
	return err
}
