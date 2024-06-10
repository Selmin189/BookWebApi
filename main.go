package main

import (
	"BookWebApi/db"
	"BookWebApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Response map[string]any

func main() {

	db.Init()
	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {
		result, err := models.GetAllBooks()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(200, Response{
			"message": "All books in the database",
			"books":   result,
		})
	})

	app.GET("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		book, err := models.GetBookById(int64(id))
		if err != nil {
			context.JSON(404, Response{
				"message": "Book not found",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book found",
			"book":    book,
		})
	})

	app.POST("/books", func(context *gin.Context) {
		var bookObject models.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		err = bookObject.Save()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book created successfully",
			"object":  bookObject,
		})
	})

	app.PUT("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		var bookObject models.Book
		err = context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		bookObject.Id = int64(id)
		err = models.UpdateBook(bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot update book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book updated successfully",
			"object":  bookObject,
		})
	})

	app.DELETE("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		err = models.DeleteBook(int64(id))
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot delete book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book deleted successfully",
		})
	})

	err := app.Run(":8088")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}
