package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/shahincsejnu/golang-testing-using-ginkgo/intro-to-BDD-in-go-using-Ginkgo/official-ginkgo/bootstrapping-a-suite/books"
)

var _ = Describe("Book", func() {
	It("can be loaded from JSON", func() {
		book := NewBook("Nothing", "No one", 2020)

		Expect(book.Title).To(Equal("Nothing"))
		Expect(book.Author).To(Equal("No one"))
		Expect(book.Pages).To(Equal(2020))
	})

	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Ginkgo 101",
			Author: "Sahadat Hossain",
			Pages:  302,
		}

		shortBook = Book{
			Title:  "Gomega 101",
			Author: "Sahadat Hossain",
			Pages:  22,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
