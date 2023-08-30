package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/asfand687/go-basic-api/models"
)

// Book Input Struct
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// usually the variable used for context is c
// i am using res so it looks familiar to express
// it acts as both req and res from express
// GET /books
// Get all books
func FindBooks(res *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	res.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(res *gin.Context) {
	// Validate input from the body
	var input CreateBookInput
	// checking validation error
	if err := res.ShouldBindJSON(&input); err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	res.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(res *gin.Context) { // Get model if exist
	var book models.Book

	// checks for error and if there is no error, it saves the info in the variable named book
	if err := models.DB.Where("id = ?", res.Param("id")).First(&book).Error; err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	res.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// update a book

func UpdateBook(res *gin.Context) {
	var book models.Book

	// checks for error and if there is no error, it saves the info in the variable named book
	if err := models.DB.Where("id = ?", res.Param("id")).First(&book).Error; err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input from the body
	var input UpdateBookInput

	// checking validation error
	if err := res.ShouldBindJSON(&input); err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// updating the model with new body
	models.DB.Model(&book).Updates(input)

	res.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(res *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", res.Param("id")).First(&book).Error; err != nil {
		res.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	res.JSON(http.StatusOK, gin.H{"data": true})
}
