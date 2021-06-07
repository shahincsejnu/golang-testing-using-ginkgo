package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (val *Book) CategoryByLength() string {
	if val.Pages > 200 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func Start() {
	fmt.Println("Demo testing...")
}
