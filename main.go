package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/kteb/test_pop/models"
)

func main() {
	db, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}
	// author := &models.Author{Name: "Auth1"}
	// db.Save(author)
	// author2 := &models.Author{Name: "Auth2"}
	// db.Save(author2)
	// book := &models.Book{Title: "Title1", AuthorID: author.ID}
	// db.Save(book)
	// book2 := &models.Book{Title: "Title2"}
	// db.Save(book2)
	authors := &models.Authors{}
	err = db.Eager().All(authors)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(authors)
	books := models.Books{}
	err = db.Eager("Author").All(&books)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(books[1])

}
