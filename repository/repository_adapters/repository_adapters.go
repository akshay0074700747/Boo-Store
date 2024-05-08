package repositoryadapters

import (
	"github.com/akshay0074700747/book_store/configurations"
	"github.com/akshay0074700747/book_store/entities"
)

type RepositoryAdapter struct {
	UsersDataPath      string
	UserBooksDataPath  string
	AdminBooksDataPath string
}

func NewRepositoryAdapter(config configurations.Configurations) *RepositoryAdapter {
	return &RepositoryAdapter{
		UsersDataPath:      config.UserDataPath,
		UserBooksDataPath:  config.UserBooksDataPath,
		AdminBooksDataPath: config.AdminBooksDataPath,
	}
}

func (repo *RepositoryAdapter) LoginUser(user entities.User) (entities.User, error) {

}
