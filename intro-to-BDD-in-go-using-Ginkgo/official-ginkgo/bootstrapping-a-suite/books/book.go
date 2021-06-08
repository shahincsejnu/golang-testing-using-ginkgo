package books

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 200 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func NewBookFromJSON(data []byte) Book {
	var book Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		return Book{}
	}
	fmt.Println(book.Title)
	return book
}

func NewBookFromJSON2(data []byte) (Book, error) {
	var book Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b *Book) AuthorLastName() string {
	name := strings.Split(b.Author, " ")
	return name[len(name)-1]
}

func DoSomething() bool {
	return true
}

func NewBook(title, author string, pages int) Book {
	return Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func Start() {
	fmt.Println("Demo testing...")
}
