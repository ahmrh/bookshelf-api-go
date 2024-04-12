package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/ahmrh/bookshelf-api-go/utils"
)

type Book struct{
	ID        	string `json:"id"`
	Name      	string `json:"name"`
	Year  			int64 `json:"year"`
	Author    	string `json:"author"`
	Summary  		string `json:"summary"`
	Publisher  	string  `json:"publisher"`
	PageCount  	int64   `json:"pageCount"`
  ReadPage    int64  `json:"readPage"`  
  Finished    bool   `json:"finished"`  
  Reading     bool   `json:"reading"`  
  InsertedAt  string `json:"insertedAt"`
  UpdatedAt   string `json:"updatedAt"`
}

type BookFilter struct {
	Name			*string
	Finished	*bool
	Reading		*bool
}

var books = [] Book {}


// handle all thing from database or data source
func GetBookByID(id string) *Book{
	for _, book := range books {
		if(book.ID == id){
			return &book
		}
	}
	return nil
}

func GetAllBooks() [] Book {
  // Create a copy of the slice to avoid modification from outside
	
  return append([]Book{}, books...) 
}

func GetBooks(filter BookFilter) []Book {
  filteredBooks := []Book{}

  for _, book := range books {
		match := true

    if filter.Name != nil && !strings.Contains(strings.ToLower(book.Name), strings.ToLower(*filter.Name)) {
      match = false 
    }
    if filter.Finished != nil && *filter.Finished != book.Finished {
      match = false 
    }
    if filter.Reading != nil && *filter.Reading != book.Reading {
      match = false 
    }

    if match {
      filteredBooks = append(filteredBooks, book)
    }
  }
  return filteredBooks
}

func AddBook(book Book) (Book, error) {

	if(book.Name == "") {
		return Book{}, fmt.Errorf("Gagal menambahkan buku. Mohon isi nama buku")
	}

	if(book.ReadPage > book.PageCount) {
		return Book{},  fmt.Errorf("Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount")
	}

	book.ID = utils.GenerateId()

	book.Finished = book.PageCount == book.ReadPage
	book.InsertedAt = time.Now().Format(time.RFC3339)
	book.UpdatedAt = book.InsertedAt

	books = append(books, book)


	return book, nil
}

func EditBookByID(id string, data Book) (*Book, error){

	if(data.Name == "") {
		return &Book{}, fmt.Errorf("Gagal memperbarui buku. Mohon isi nama buku")
	}

	if(data.ReadPage > data.PageCount) {
		return &Book{}, fmt.Errorf("Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount")
	}
	

	for i, book := range books {
		if(book.ID == id){
			
			books[i].Name = data.Name
			books[i].Year = data.Year
			books[i].Author = data.Author
			books[i].Summary = data.Summary
			books[i].Publisher = data.Publisher
			books[i].PageCount = data.PageCount
			books[i].ReadPage = data.ReadPage
			books[i].Reading = data.Reading


			return &book, nil
		}
	}

	return nil, fmt.Errorf("Gagal memperbarui buku. Id tidak ditemukan")

}

func DeleteBookByID(id string) error {

	for i, book := range books {
		if book.ID == id{

			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Buku gagal dihapus. Id tidak ditemukan")

}
