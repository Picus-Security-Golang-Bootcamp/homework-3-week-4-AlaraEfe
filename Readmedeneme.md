# Homework 3| Week 4

This program contains a list of books in PostgreSQL database.

Program have a Book Model, implement the struct to the database and Insert the data of .csv file to database.

In that program db queries and GORM queries had been used to build functions.

## The Functions and Usage

### ListAllBooksWithRawSQL()

``` 
ListAllBooksWithRawSQL()  
``` 
* ListAllBooksWithRawSQL() function to list all the books with details in the Book_Archive database with RawSQL

### SearchByNameWithRawSQL("bookName")

``` 
SearchByNameWithRawSQL("bookName") 
``` 
* SearchByNameWithRawSQL function to search and gets the book or books with all details in the book list of Book_Archive database with given book name or the word included in the book name

### SoftDeleteGormSQL("bookID")

``` 
SoftDeleteGormSQL("bookID")  
``` 
* SoftDeleteGormSQL function to soft delete the book from database with given book ID in the book list of Book_Archive database with GormSQL. The deleted file can be seen with using GetByID function

### BuyBookWithRawSQL("bookID", quantity)

``` 
BuyBookWithRawSQL("bookID", quantity)
``` 
* BuyBookWithRawSQL function to buy the book and update the current book stock information in the database. Also, it calculates the total cost of the book.

### GetByID("bookID")

``` 
GetByID("bookID")
``` 
* GetByID function to gets the book by book ID. If there is soft deleted records, book also returns with the message "That book was deleted from Book Archive database"

### FindByName("bookName")

``` 
FindByName("bookName")
``` 
* FindByName function finds the the book of Book_Archive database with the given book name

### GetBooksWithAuthorID("authorID")

``` 
GetBooksWithAuthorID("authorID") 
``` 
* GetBooksWithAuthorID function to get books with given Author ID

### GetAuthorWithBooks("bookName")

``` 
GetAuthorWithBooks("bookName") 
``` 
* GetAuthorWithBooks function to get Author Name of the book with given book name