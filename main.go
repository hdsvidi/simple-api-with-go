package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID 			string 	`json:"id"`
	Title 		string	`json:"title"`
	Author 		string	`json:"author"`
	Quantity 	int		`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The lord of the Ring", Author: "Tolkien", Quantity: 5},
	{ID: "2", Title: "To Kill a Mocking Jay", Author: "Harper lee", Quantity: 4},
	{ID: "3", Title: "Twilight sage : eclipse", Author: "Stephany mayer", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createdBook(c * gin.Context) {
	var payload book
	
	if err := c.BindJSON(&payload); err != nil {
		return
	}

	books = append(books, payload)
	c.IndentedJSON(http.StatusCreated, payload)
}

func getBookById(id string) (*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book is not found")
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book is not found"})
		return 
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkout(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "missing id query"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book is not found"})
		return 
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "book is not availabe"})
	}

	book.Quantity -= 1;
	c.IndentedJSON(http.StatusOK, book)
}

func checkin(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "missing id query"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book is not found"})
		return 
	}

	book.Quantity += 1;
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createdBook)
	router.PATCH("/checkout", checkout)
	router.PATCH("/checkin", checkin)
	router.Run()
}

 