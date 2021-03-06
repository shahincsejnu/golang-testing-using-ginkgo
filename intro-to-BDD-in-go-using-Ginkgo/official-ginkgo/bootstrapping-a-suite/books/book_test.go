package books_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/shahincsejnu/golang-testing-using-ginkgo/intro-to-BDD-in-go-using-Ginkgo/official-ginkgo/bootstrapping-a-suite/books"
)

var _ = Describe("Temp", func() {
	var book Book
	BeforeEach(func() {
		book = NewBookFromJSON([]byte(`{
            "title":"first one",
            "author":"Victor Hugo",
            "pages":2783
        }`))
	})

	It("", func() {
		Expect(book.Author).To(Equal("Victor Hugo"))
	})

	Context("first level", func() {
		BeforeEach(func() {
			book = NewBookFromJSON([]byte(`{
				"title":"second one",
				"author":"Victor Hugo",
				"pages":2783
			}`))
		})

		It("", func() {
			Expect(book.Author).To(Equal("Victor Hugo"))
		})

		Context("second level", func() {
			BeforeEach(func() {
				book = NewBookFromJSON([]byte(`{
					"title":"third one",
					"author":"Victor Hugo",
					"pages":2783
				}`))
			})

			It("", func() {
				Expect(book.Author).To(Equal("Victor Hugo"))
			})
		})

		Context("second level2", func() {
			BeforeEach(func() {
				book = NewBookFromJSON([]byte(`{
					"title":"fourth",
					"author":"Victor Hugo",
					"pages":2783
				}`))
			})

			It("", func() {
				Expect(book.Author).To(Equal("Victor Hugo"))
			})
		})
	})

	Context("first level2", func() {
		BeforeEach(func() {
			book = NewBookFromJSON([]byte(`{
				"title":"fifth",
				"author":"Victor Hugo",
				"pages":2783
			}`))
		})

		It("", func() {
			Expect(book.Author).To(Equal("Victor Hugo"))
		})
	})
})

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
	)

	BeforeEach(func() {
		book, err = NewBookFromJSON2([]byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`))
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(2783))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON2([]byte(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":2783oops
                }`))
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
		json []byte
	)

	BeforeEach(func() {
		json = []byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`)
	})

	JustBeforeEach(func() {
		book, err = NewBookFromJSON2(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(2783))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				json = []byte(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":2783oops
                }`)
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var _ = Describe("Book", func() {
	var book Book

	BeforeEach(func() {
		book = NewBookFromJSON([]byte(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":2783
        }`))
	})

	It("can be loaded from JSON", func() {
		Expect(book.Title).To(Equal("Les Miserables"))
		Expect(book.Author).To(Equal("Victor Hugo"))
		Expect(book.Pages).To(Equal(2783))
		book.Author = "Oka Oka"
		fmt.Println(book)
	})

	It("can extract the author's last name", func() {
		fmt.Println(book)
		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})
})

var _ = Describe("Book", func() {
	It("can be loaded from JSON", func() {
		book := NewBook("Nothing", "No one", 2020)
		fmt.Println(CurrentGinkgoTestDescription())
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

		//FIt("can be loaded from JSON", func() {
		//	book := NewBookFromJSON([]byte(`{
		//    "title":"Les Miserables",
		//    "author":"Victor Hugo",
		//    "pages":2783
		//}`))
		//
		//	Expect(book.Title).To(Equal("Les Miserables"))
		//	Expect(book.Author).To(Equal("Victor Hugo"))
		//	Expect(book.Pages).To(Equal(2783))
		//})
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
