package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	Book1.title = "Go Programmin"
	Book1.author = "Mursalin"
	Book1.subject = "Golang "
	Book1.book_id = 1123

	Book2.title = "My Life"
	Book2.author = "Kabir"
	Book2.subject = "Life "
	Book2.book_id = 3344

	printBook(Book1)

	printBook(Book2)

}
func printBook(book Books) {
	fmt.Printf("Book Title: %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)

}
