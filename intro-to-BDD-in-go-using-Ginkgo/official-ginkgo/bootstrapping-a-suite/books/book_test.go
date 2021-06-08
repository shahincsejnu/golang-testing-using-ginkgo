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

	var _ = Describe("Book", func() {
		It("can be loaded from JSON", func() {
			book := NewBookFromJSON([]byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`))

			Expect(book.Title).To(Equal("Les Miserables"))
			Expect(book.Author).To(Equal("Victor Hugo"))
			Expect(book.Pages).To(Equal(2783))
		})

		PIt("can be loaded from JSON", func() {
			book := NewBookFromJSON([]byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`))

			Expect(book.Title).To(Equal("Les Miserables"))
			Expect(book.Author).To(Equal("Victor Hugo"))
			Expect(book.Pages).To(Equal(2783))
		})

		XIt("can be loaded from JSON", func() {
			book := NewBookFromJSON([]byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`))

			Expect(book.Title).To(Equal("Les Miserables"))
			Expect(book.Author).To(Equal("Victor Hugo"))
			Expect(book.Pages).To(Equal(2783))
		})
	})

	done := make(chan bool)

	It("panics in a goroutine", func() {
		go func() {
			defer GinkgoRecover()

			Expect(DoSomething()).Should(BeTrue())
			close(done)
		}()
		//_, ok := <- done
		//Expect(ok).To(BeFalse())
		Eventually(done).Should(BeClosed())
	})

	Describe("The foobar service", func() {
		Context("when calling Foo()", func() {
			Context("when no ID is provided", func() {
				Specify("an ErrNoID error is returned", func() {

				})
			})
		})
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
