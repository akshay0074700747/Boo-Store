package controllers

import (
	"fmt"

	"github.com/akshay0074700747/book_store/web/handlers"
	"github.com/gin-gonic/gin"
)

type BookStoreController struct {
	controllers *gin.Engine
}

func NewBookStoreController(userHandler *handlers.UserHandler, booksHandler *handlers.BookHandlers) *BookStoreController {

	engine := gin.New()

	engine.Use(gin.Logger())

	engine.POST("/user/login", userhandler.UserLogin)

	return &BookStoreController{
		controllers: engine,
	}
}

func (ctrl *BookStoreController) Start(port string) {

	fmt.Printf("Server is Starting on port %s ...", port)

	//starting up the server
	ctrl.controllers.Run(port)
}
