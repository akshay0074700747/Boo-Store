package repositoryports

import "github.com/akshay0074700747/book_store/entities"

type RepositoryPort interface {
	LoginUser(user entities.User) (entities.User,error)
}