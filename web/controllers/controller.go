package controllers

import (
	"fmt"

	"github.com/akshay0074700747/book_store/web/handlers"
	"github.com/akshay0074700747/book_store/web/middlewares"
	"github.com/gin-gonic/gin"
)

type BookStoreController struct {
	controllers *gin.Engine
}

func NewBookStoreController(handlers *handlers.Handler, middleware *middlewares.Middleware) *BookStoreController {

	//initialising the gin server
	engine := gin.New()

	//setting the logger for logging out the events in the server
	engine.Use(gin.Logger())

	engine.POST("/login", handlers.Login)

	user := engine.Group("/user")
	user.Use(middleware.GlobalMiddleware())
	{
		user.GET("/home", handlers.GetHome)

		admin := user.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/addBook", handlers.AddBook)
			admin.DELETE("/deleteBook/:bookName", handlers.DeleteBook)
		}
	}

	return &BookStoreController{
		controllers: engine,
	}
}

//starts the server in the specied port
func (ctrl *BookStoreController) Start(port string) {

	fmt.Printf("Server is Starting on port %s ...", port)

	//starting up the server
	ctrl.controllers.Run(port)
}
