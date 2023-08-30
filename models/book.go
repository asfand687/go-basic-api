package models

// Defining the models using struct
// struct is similar to a database object for defining the data
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
