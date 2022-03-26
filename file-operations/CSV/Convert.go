package csv_utils

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/AlaraEfe/file-operations/file-operations/models"
)

func JSONToCSV(source string, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var books models.BooksSlice
	err = json.NewDecoder(sourceFile).Decode(&books)
	if err != nil {
		return err
	}

	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := csv.NewWriter(output)
	headers := []string{"Book_ID", "Book_Name", "Book_Page_Count", "Book_Stock", "Book_Price", "Book_Stock_Code", "Book_ISBN", "Book_Author_ID", "BookAuthorName"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, book := range books {
		var row []string
		row = append(row, book.BookID, book.BookName, book.BookPageCount, book.BookStock, book.BookPrice, book.BookStockCode, book.BookISBN, book.BookAuthorID, book.BookAuthorName)
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	log.Println("JSON to CSV process completed")
	return nil
}
