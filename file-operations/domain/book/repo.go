package book

import (
	"errors"
	"fmt"

	"github.com/AlaraEfe/file-operations/file-operations/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertData(books models.BooksSlice) {
	for _, book := range books {
		b.db.Where(Book{BookID: book.BookID}).
			Attrs(Book{BookID: book.BookID, BookName: book.BookName}).FirstOrCreate(&book)
	}

}

func (b *BookRepository) ListAllBooksWithRawSQL() []Book {
	var books []Book
	b.db.Raw("SELECT * FROM books").Scan(&books)

	return books
}

func (b *BookRepository) SearchByNameWithRawSQL(name string) []Book {
	var books []Book
	b.db.Raw("SELECT * FROM books WHERE book_name ILIKE ?", "%"+name+"%").Find(&books)

	return books
}

func (b *BookRepository) SoftDeleteGormSQL(ID string) *Book {
	var book Book
	b.db.Where("book_id = ?", ID).Delete(&book)

	return &book
}

func (b *BookRepository) BuyBookWithRawSQL(ID string, quantitiy int) (*Book, error) {
	var book Book

	result := b.db.Raw("SELECT * FROM books WHERE book_id LIKE ?", "%"+ID+"%").First(&book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	if book.BookStock >= quantitiy {

		b.db.Model(&book).UpdateColumn("book_stock", gorm.Expr("book_stock - ?", quantitiy))

	} else {
		fmt.Println("There is not enough stock of that book")
		return nil, result.Error
	}

	bookPrice := book.BookPrice

	totalBookCost := bookPrice * float64(quantitiy)

	s := fmt.Sprintf("The total cost of %d book/books named '%s' is: %2.f\n ", quantitiy, book.BookName, totalBookCost)
	print(s)

	return &book, nil
}

func (b *BookRepository) GetByID(id string) (*Book, error) {
	var book Book

	result := b.db.Where("book_id = ?", id).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		softDeletedRecords := b.db.Unscoped().Where("book_id = ?", id).First(&book)
		if errors.Is(softDeletedRecords.Error, gorm.ErrRecordNotFound) {

			return nil, result.Error

		}

		fmt.Println("That book was deleted from Book Archive database")

		return &book, nil

	}

	return &book, nil
}

func (b *BookRepository) FindByName(bookname string) (*Book, error) {
	var book Book
	result := b.db.Where("book_name ILIKE ?", "%"+bookname+"%").First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	fmt.Println(book.BookAuthorID)
	fmt.Println(book.BookID)

	return &book, nil
}

func (b *BookRepository) GetBooksWithAuthorID(authorID string) []Book {
	var books []Book
	b.db.Where("book_author_id = ?", authorID).Find(&books)

	return books
}

func (b *BookRepository) GetAuthorWithBooks(bookName string) (*Book, error) {
	var book Book
	result := b.db.Where("book_name ILIKE ?", "%"+bookName+"%").First(&book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	s := fmt.Sprintf("The author of %s is %s.\n ", book.BookName, book.BookAuthorName)
	print(s)

	return &book, nil
}
