package repositoryports

import "github.com/akshay0074700747/book_store/entities"

//repositoryport is the abstration interface for achieving loosely coupling between dependencies
type RepositoryPort interface {
	LoginUser(user entities.User) (entities.User, error)
	GetBooks(isAdmin bool) ([]entities.Book, error)
	AddBook(book entities.Book) error
	DeleteBook(bookName string) error
}
