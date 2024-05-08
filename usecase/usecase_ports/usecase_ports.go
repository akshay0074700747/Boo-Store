package usecaseports

import "github.com/akshay0074700747/book_store/entities"

type UsecasePort interface {
	LoginUser(user entities.User) (entities.User,error)
}