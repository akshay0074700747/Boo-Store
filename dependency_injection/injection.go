package dependencyinjection

import "github.com/akshay0074700747/book_store/configurations"

func Initialize(cfg configurations.Configurations) *services.UserEngine {

	db := db.ConnectDB(cfg)
	adapter := adapters.NewUserAdapter(db)
	usecase := usecases.NewUserUsecases(adapter)
	server := services.NewUserServiceServer(usecase, "auth-service:50004")

	return services.NewUserEngine(server)
}