package usecaseadapters

import (
	"errors"

	customerrormessages "github.com/akshay0074700747/book_store/custom_errormessages"
	"github.com/akshay0074700747/book_store/entities"
	"github.com/akshay0074700747/book_store/helpers"
	repositoryports "github.com/akshay0074700747/book_store/repository/repository_ports"
)

//usecaseadapter implements the usecaseport interface
type UsecaseAdapter struct {
	Repo repositoryports.RepositoryPort
}

func NewUsecaseAdapter(repo repositoryports.RepositoryPort) *UsecaseAdapter {
	return &UsecaseAdapter{
		Repo: repo,
	}
}

func (usecase *UsecaseAdapter) LoginUser(user entities.User) (entities.User, error) {

	result, err := usecase.Repo.LoginUser(user)
	if err != nil {
		return entities.User{}, err
	}
	// setting the password to empty in the result 
	result.Password = ""

	return result, nil
}

func (usecase *UsecaseAdapter) GetBooks(isAdmin bool) ([]entities.Book, error) {

	var result []entities.Book

	//checking whether the home request comes from user or admin
	if isAdmin {
		adminBooks, err := usecase.Repo.GetBooks(isAdmin)
		if err != nil {
			return nil, err
		}

		userBooks, err := usecase.Repo.GetBooks(!isAdmin)
		if err != nil {
			return nil, err
		}

		result = append(result, adminBooks...)
		result = append(result, userBooks...)
	} else {

		userBooks, err := usecase.Repo.GetBooks(isAdmin)
		if err != nil {
			return nil, err
		}

		result = append(result, userBooks...)
	}

	return result, nil
}

func (usecase *UsecaseAdapter) AddBook(book entities.Book) error {

	//validating book name,author name and publish date
	if !helpers.ValidateString(book.BookName) {
		return errors.New(customerrormessages.BooknameValidationError)
	}

	if !helpers.ValidateString(book.Author) {
		return errors.New(customerrormessages.AuthornameValidationError)
	}

	if !helpers.ValidateNumber(book.PublicationYear) {
		return errors.New(customerrormessages.YearValidationError)
	}

	return usecase.Repo.AddBook(book)
}

func (usecase *UsecaseAdapter) DeleteBook(bookName string) error {

	if bookName == "" {
		return errors.New("the bookname cannot be empty")
	}

	return usecase.Repo.DeleteBook(bookName)
}
