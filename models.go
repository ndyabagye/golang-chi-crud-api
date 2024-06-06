package main

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var books = []*Book{
	{
		ID:               "1",
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

// func listBooks() []*Book {
// 	return books
// }

// func getBook(id string) *Book {
// 	for _, book := range books {
// 		if book.ID == id {
// 			return book
// 		}
// 	}
// 	return nil
// }

// func storeBook(book Book) {
// 	books = append(books, &book)
// }

// func updateBook(id string, bookUpdate Book) *Book {
// 	for i, book := range books {
// 		if book.ID == id {
// 			books[i] = &bookUpdate
// 			return book
// 		}
// 	}
// 	return nil
// }

// func deleteBook(id string) *Book {
// 	for i, book := range books {
// 		if book.ID == id {
// 			books = append(books[:i], (books)[i+1:]...)
// 			return &Book{}
// 		}
// 	}
// 	return nil
// }

type BookStorage interface {
	List() []*Book
	Get(string) *Book
	Update(string, Book) *Book
	Create(Book)
	Delete(string) *Book
}

type BookStore struct {
}

func (b BookStore) Get(id string) *Book {
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func (b BookStore) List() []*Book {
	return books
}

func (b BookStore) Create(book Book) {
	books = append(books, &book)
}

func (b BookStore) Delete(id string) *Book {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], (books)[i+1:]...)
			return &Book{}
		}
	}
	return nil
}

func (b BookStore) Update(id string, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = &bookUpdate
			return book
		}
	}
	return nil
}
