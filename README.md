# Basic CRUD Book Application with Gin and MySQL

This is a simple CRUD (Create, Read, Update, Delete) application for managing books using the Gin web framework and MySQL database.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Go programming language is installed. [Install Go](https://golang.org/doc/install)
- Gin framework is installed. Run `go get -u github.com/gin-gonic/gin`
- MySQL database is installed and running.

## Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/asfand687/go-basic-api
   go-basic-api

2. Install dependencies
  ```bash
  go mod tidy

3. Access the application in your web browser at http://localhost:8080

**Endpoints**
* GET /users Get dummy users
* GET /books: Get a list of all books.
* GET /book/:id: Get details of a specific book by ID.
* POST /books: Add a new book. Provide JSON payload with title and author fields.
* PATCH /book/:id: Update details of a specific book by ID. Provide JSON payload with fields to update.
* DELETE /book/:id: Delete a specific book by ID.
