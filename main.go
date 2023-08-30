package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/asfand687/go-basic-api/controllers"
	"www.github.com/asfand687/go-basic-api/models"
)

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// This is the dummy data if you dont want to use any db. Just for the get endpoint
var users = []user{
	{Id: 546, Username: "John"},
	{Id: 894, Username: "Mary"},
	{Id: 326, Username: "Jane"},
}

// TODO read more about gin.Context
func getUsers(res *gin.Context) {
	res.IndentedJSON(http.StatusOK, users)
}

// Init function, just like the top of the main.js file in express
func init() {
	models.ConnectDatabase()
}

// Main function where s**t happens
func main() {
	// const router = express.router()
	router := gin.Default()

	// new `GET /users` route associated with our `getUsers` function
	// app.get("/", getUsers)
	router.GET("/users", getUsers)

	// Find Books
	router.GET("/books", controllers.FindBooks)

	// Create Book
	router.POST("/books", controllers.CreateBook)

	// Get Single Book
	router.GET("/book/:id", controllers.FindBook)

	// Update Book
	router.PATCH("/book/:id", controllers.UpdateBook)

	// Delete Book
	router.DELETE("/book/:id", controllers.DeleteBook)

	// app.listen(8000, () => {console.log("Server Running on 8000")})
	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
