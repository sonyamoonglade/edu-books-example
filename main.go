package main

import (
	"fmt"
	"log"

	book_storage "sonya/books_storage"
)

var (
	b1 = &book_storage.Book{
		Title:  "English",
		Pages:  125,
		Author: "John Tomer",
		Grade:  11,
	}
	b2 = &book_storage.Book{
		Title:  "Math",
		Pages:  125,
		Author: "John Tomer",
		Grade:  11,
	}

	b3 = &book_storage.Book{
		Title:  "Algebra",
		Pages:  123123,
		Author: "Marinka",
		Grade:  11,
	}
)

func main() {
	// for _, book := range books{
	// fmt.Printf("name: %s, pages: %d
	//}

	bs := book_storage.NewBookStorage(3)

	err := bs.AddBook(b1)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = bs.AddBook(b2)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = bs.AddBook(b3)
	if err != nil {
		log.Fatal(err.Error())
	}

	books11 := bs.BooksByGrade(11)
	for _, b := range books11 {
		fmt.Printf("title: %s, pages: %d, author: %s\n", b.Title, b.Pages, b.Author)
	}

}
