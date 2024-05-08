package dependencyinjection

import (
	"github.com/akshay0074700747/book_store/configurations"
	repositoryadapters "github.com/akshay0074700747/book_store/repository/repository_adapters"
	usecaseadapters "github.com/akshay0074700747/book_store/usecase/usecase_adapters"
	"github.com/akshay0074700747/book_store/web/controllers"
	"github.com/akshay0074700747/book_store/web/handlers"
)

func InjectDependencies(cfg configurations.Configurations) *controllers.BookStoreController {

	repo := repositoryadapters.NewRepositoryAdapter(cfg)
	usecase := usecaseadapters.NewUsecaseAdapter(repo)
	handlers := handlers.NewHandler(usecase, cfg.Secret)
	return controllers.NewBookStoreController(handlers)
}
