package controllers

import (
	"log"
	"net/http"
	"github.com/ahmrh/bookshelf-api-go/models"
	"github.com/ahmrh/bookshelf-api-go/utils"
	"github.com/gin-gonic/gin"
)

type BookController struct{}

func (b BookController) Create(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		log.Printf("Error occured: %v", err)
		return
	}

	if newBook, err := models.AddBook(newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Buku berhasil ditambahkan",
			"data":    gin.H{"bookId": newBook.ID},
		})
	}
}

func (b BookController) Read(c *gin.Context) {

  bookId := c.Param("id")

  if bookId != "" {
		if book := models.GetBookByID(bookId); book == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": "Buku tidak ditemukan",
			})
		} else {

			c.IndentedJSON(http.StatusOK, gin.H{
				"status":  "success",
				"data":    gin.H{"book": book},
			})
		}
		
		
 } else {
		var namePtr *string
		if nameQuery := c.Query("name"); nameQuery == "" {
			namePtr = nil
		} else {
			namePtr = &nameQuery
		}
		reading, _ := utils.StringToBool(c.Query("reading"))
		finished, _ := utils.StringToBool(c.Query("finished"))

		filteredBooks := models.GetBooks( models.BookFilter{Name: namePtr, Reading: reading, Finished: finished })

		books := make([]gin.H, 0, len(filteredBooks))

		for _, book := range filteredBooks {
			books = append(books, gin.H{
				"id"			: book.ID,
				"name"		:	book.Name,
				"publisher": book.Publisher,
			})
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "success",
			"data":    gin.H{"books": books},
		})
 }

}

func (b BookController) Update(c *gin.Context) {

	var bookData models.Book

	if err := c.BindJSON(&bookData); err != nil {
		log.Printf("Error occured: %v", err)
		return
	}

  bookId := c.Param("id")

  if bookId != "" {
		if book, err := models.EditBookByID(bookId, bookData); book == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {

			if(err != nil) {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"status":  "fail",
					"message": err.Error(),
				})
			} else {

				c.IndentedJSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "Buku berhasil diperbarui",
				})
			}
		}

	
 }
}


func (b BookController) Delete(c *gin.Context) {
	
  bookId := c.Param("id")

  if bookId != "" {
		if err := models.DeleteBookByID(bookId); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Buku berhasil dihapus",
			})
		}

	
 }
}