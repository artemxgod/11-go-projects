package model

import (
	"log"

	"github.com/artemxgod/11-go-projects/book-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.DB().Ping(); err != nil {
		log.Fatal("err db", err)
	}
	db.AutoMigrate(Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var res Book
	db := db.Where("ID=?", ID).Find(&res)
	return &res, db
}

func DeleteBook(ID int64) Book {
	var res Book
	db.Where("ID=?", ID).Delete(&res)
	return res
}