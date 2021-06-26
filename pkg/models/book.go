package models

import (
	"github.com/KaT0819/sample-go-gorilla-api/pkg/config/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) Create() *Book {

	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAll() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) Get(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func (b *Book) Update(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db

}

func (b *Book) Delete(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book

}