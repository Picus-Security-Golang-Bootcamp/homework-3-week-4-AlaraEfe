package models

type Book struct {
	BookID         string `json:"BookID"`
	BookName       string `json:"BookName"`
	BookPageCount  string `json:"BookPageCount"`
	BookStock      string `json:"BookStock"`
	BookPrice      string `json:"BookPrice"`
	BookStockCode  string `json:"BookStockCode"`
	BookISBN       string `json:"BookISBN"`
	BookAuthorID   string `json:"BookAuthorID"`
	BookAuthorName string `json:"BookAuthorName"`
}

type BooksSlice []Book
