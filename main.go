package main

import (
	"fmt"
	"log"
	"os"

	csv_utils "github.com/AlaraEfe/file-operations/file-operations/CSV"
	postgres "github.com/AlaraEfe/file-operations/file-operations/common/db"
	"github.com/AlaraEfe/file-operations/file-operations/domain/book"
	"github.com/joho/godotenv"
)

var filename = "booksandauthors.csv"

func main() {

	//convert JSON to CSV
	/*err := csv_utils.JSONToCSV("booksandauthors.json", "booksandauthors.csv")
	if err != nil {
		log.Fatal(err)
	}*/

	//Set environment variables to PostGreSQL
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println("xxx")

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	//Book repository
	bookRepo := book.NewBookRepository(db)

	//Implement the struct in database
	bookRepo.Migrations()

	//Read the booksandauthors.csv file
	books, err := csv_utils.ReadCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	//Insert the data of booksandauthor.csv file to database
	bookRepo.InsertData(books)

	//ListAllBooksWithRawSQL() function to list all the books with details in the Book_Archive database with RawSQL
	//fmt.Println(bookRepo.ListAllBooksWithRawSQL())

	//SearchByNameWithRawSQL function to search and gets the book or books with all details in the book list of Book_Archive database with given book name or the word included in the book name
	//fmt.Println(bookRepo.SearchByNameWithRawSQL("cRimE"))

	//SoftDeleteGormSQL function to soft delete the book from database with given book ID in the book list of Book_Archive database with GormSQL. The deleted file can be seen with using GetByID function
	//bookRepo.SoftDeleteGormSQL("3")

	//BuyBookWithRawSQL function to buy the book and update the current book stock information in the database. Also, it calculates the total cost of the book.
	//bookRepo.BuyBookWithRawSQL("6", 3)

	//GetByID function to gets the book by book ID. If there is soft deleted records, book also returns with the message "That book was deleted from Book Archive database"
	//fmt.Println(bookRepo.GetByID("3"))

	//FindByName function finds the the book of Book_Archive database with the given book name
	//fmt.Println(bookRepo.FindByName("Guns, Germs, and Steel"))

	//GetBooksWithAuthorID function to get books with given Author ID
	//fmt.Println(bookRepo.GetBooksWithAuthorID("412"))

	//GetAuthorWithBooks function to get Author Name of the book with given book name
	//fmt.Println(bookRepo.GetAuthorWithBooks("Crime and Punishment"))

}
