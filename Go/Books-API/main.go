package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book struct (Model)
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// List of books
var books = []book{
	{
		ID:       "1",
		Title:    "To Kill a Mockingbird",
		Author:   "Harper Lee",
		Quantity: 10,
	},
	{
		ID:       "2",
		Title:    "1984",
		Author:   "George Orwell",
		Quantity: 5,
	},
	{
		ID:       "3",
		Title:    "The Great Gatsby",
		Author:   "F. Scott Fitzgerald",
		Quantity: 8,
	},
	{
		ID:       "4",
		Title:    "Pride and Prejudice",
		Author:   "Jane Austen",
		Quantity: 12,
	},
	{
		ID:       "5",
		Title:    "The Catcher in the Rye",
		Author:   "J.D. Salinger",
		Quantity: 7,
	},
	{
		ID:       "6",
		Title:    "Brave New World",
		Author:   "Aldous Huxley",
		Quantity: 9,
	},
	{
		ID:       "7",
		Title:    "Moby-Dick",
		Author:   "Herman Melville",
		Quantity: 3,
	},
	{
		ID:       "8",
		Title:    "The Lord of the Rings",
		Author:   "J.R.R. Tolkien",
		Quantity: 15,
	},
	{
		ID:       "9",
		Title:    "Fahrenheit 451",
		Author:   "Ray Bradbury",
		Quantity: 6,
	},
	{
		ID:       "10",
		Title:    "The Hobbit",
		Author:   "J.R.R. Tolkien",
		Quantity: 11,
	},
}

func getBooks(c *gin.Context) {
	// Handler function to get all books
	c.IndentedJSON(http.StatusOK, books) // Respond with all books in JSON format
}

func bookById(c *gin.Context) {
	// Handler function to get a book by its ID
	id := c.Param("id")          // Extract the book ID from the URL path parameter
	book, err := getBookById(id) // Call a helper function to find the book by ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, book) // Respond with the found book in JSON format
}

func checkoutBook(c *gin.Context) {
	// Handler function to checkout a book by querying its ID
	id, ok := c.GetQuery("id") // Extract the book ID from the query parameter

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter!"})
		return
	}

	book, err := getBookById(id) // Call a helper function to find the book by ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available!"})
		return
	}

	book.Quantity--                     // Decrement the book's quantity
	c.IndentedJSON(http.StatusOK, book) // Respond with the updated book in JSON format
}

func returnBook(c *gin.Context) {
	// Handler function to return a book by querying its ID
	id, ok := c.GetQuery("id") // Extract the book ID from the query parameter

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter!"})
		return
	}

	book, err := getBookById(id) // Call a helper function to find the book by ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	book.Quantity++                     // Increment the book's quantity
	c.IndentedJSON(http.StatusOK, book) // Respond with the updated book in JSON format
}

func getBookById(id string) (*book, error) {
	// Helper function to search for a book by its ID
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil // Return a pointer to the found book and no error
		}
	}

	return nil, errors.New("book not found") // Return an error if the book is not found
}

func createBooks(c *gin.Context) {
	// Handler function to add a new book
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)              // Append the new book to the list
	c.IndentedJSON(http.StatusCreated, newBook) // Respond with the new book in JSON format
}

func main() {
	fmt.Println("Hello World!")
	// Initialize the Gin router
	router := gin.Default()

	// Define routes for different API actions
	router.GET("/books", getBooks)          // Route to get all books
	router.GET("/books/:id", bookById)      // Route to get a book by ID
	router.POST("/books", createBooks)      // Route to create a new book
	router.PATCH("/checkout", checkoutBook) // Route to checkout a book
	router.PATCH("/return", returnBook)     // Route to return a book

	// Run the server on localhost:8080
	router.Run("localhost:8080")
}
