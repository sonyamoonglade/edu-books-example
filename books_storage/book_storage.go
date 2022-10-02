package book_storage

import (
	"fmt"
)

var (
	maxGrades = 11
)

type Book struct {
	Title  string
	Pages  int
	Author string
	Grade  int
	gdzURL string
}

type BookStorage struct {
	books   map[int][]*Book
	bookCap int
}

func NewBookStorage(bookCap int) *BookStorage {
	bs := &BookStorage{
		books:   make(map[int][]*Book, maxGrades),
		bookCap: bookCap,
	}
	return bs
}

func (bs *BookStorage) AddBook(b *Book) error {

	bb := bs.books[b.Grade]
	if len(bb) == bs.bookCap {
		return fmt.Errorf("no more book space. Book - title: %s", b.Title)
	}

	bs.books[b.Grade] = append(bb, b)

	return nil
}

func (bs *BookStorage) BooksByGrade(grade int) []*Book {
	return bs.books[grade]
}
