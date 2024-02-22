package models

import (
	"github.com/anhtoan1103/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// define a function that belongs to Book type with return value is an pointer to Book
func (b *Book) CreateBook() *Book {
	// check if value's primary key is blank
	// should we handle the case the primary key is blank or we should  check it in other function?
	db.NewRecord(b)
	// insert the value into database
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	// Find records that match the given condition, in case the empty condition, it returns all the records.
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
