package usecaseadapters

import (
	"github.com/akshay0074700747/book_store/entities"
	repositoryports "github.com/akshay0074700747/book_store/repository/repository_ports"
)

type UsecaseAdapter struct {
	Repo repositoryports.RepositoryPort
}

func NewUsecaseAdapter(repo repositoryports.RepositoryPort) *UsecaseAdapter {
	return &UsecaseAdapter{
		Repo: repo,
	}
}

func (usecase *UsecaseAdapter) LoginUser(user entities.User) (entities.User,error) {
	
}