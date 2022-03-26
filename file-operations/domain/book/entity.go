package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookID         string
	BookName       string
	BookPageCount  string
	BookStock      int
	BookPrice      float64
	BookStockCode  string
	BookISBN       string
	BookAuthorID   string
	BookAuthorName string
}
