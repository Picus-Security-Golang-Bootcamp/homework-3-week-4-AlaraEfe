package csv_utils

import (
	"encoding/csv"
	"os"

	"github.com/AlaraEfe/file-operations/file-operations/models"
)

func ReadCSV(filename string) (models.BooksSlice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books models.BooksSlice
	for _, line := range records[1:] {
		books = append(books, models.Book{
			BookID:         line[0],
			BookName:       line[1],
			BookPageCount:  line[2],
			BookStock:      line[3],
			BookPrice:      line[4],
			BookStockCode:  line[5],
			BookISBN:       line[6],
			BookAuthorID:   line[7],
			BookAuthorName: line[8],
		})
	}

	return books, nil
}
