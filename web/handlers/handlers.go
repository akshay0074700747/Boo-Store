package handlers

import (
	"errors"
	"net/http"

	customerrormessages "github.com/akshay0074700747/book_store/custom_errormessages"
	"github.com/akshay0074700747/book_store/entities"
	usecaseports "github.com/akshay0074700747/book_store/usecase/usecase_ports"
	jwttoken "github.com/akshay0074700747/book_store/web/jwt_token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase usecaseports.UsecasePort
	secret  string
}

//servers structured responce to the requests
type Response struct {
	StatusCode int         `json:"stastuscode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"error,omitempty"`
}

//newhandler struct takes the usecaseport interface as parameter for achieving loosely coupling between dependencies
func NewHandler(usecase usecaseports.UsecasePort, secret string) *Handler {

	return &Handler{
		secret:  secret,
		Usecase: usecase,
	}
}

func (handler *Handler) Login(c *gin.Context) {

	cookie, _ := c.Request.Cookie("Token")

	if cookie != nil {
		c.JSON(http.StatusConflict, Response{
			StatusCode: 409,
			Message:    "already logged in",
			Data:       nil,
			Errors:     nil,
		})
		return
	}

	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 400,
			Message:    "the request body is not valid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := handler.Usecase.LoginUser(user)
	if err != nil {
		if err.Error() == customerrormessages.UserNotFound {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 400,
				Message:    "the user deoesnt have an account",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 500,
			Message:    "cannot bind the request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	jwt, err := jwttoken.GenerateJwt(result.Username, result.IsAdmin, []byte(handler.secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 500,
			Message:    "couldn't generate jwt token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetCookie("Token", jwt, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, Response{
		StatusCode: 200,
		Message:    "user logged in successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (handler *Handler) GetHome(c *gin.Context) {

	value, exists := c.Get("values")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, errors.New(customerrormessages.CredentialsNotfound))
		return
	}

	valueMap, _ := value.(map[string]interface{})

	isAdmin := valueMap["isAdmin"].(bool)

	result, err := handler.Usecase.GetBooks(isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 500,
			Message:    "couldn't get list of books",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	var msg string
	if isAdmin {
		msg = "fetched books for admin successfully"
	} else {
		msg = "fetched books for user successfully"
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 200,
		Message:    msg,
		Data:       result,
		Errors:     nil,
	})
}

func (handler *Handler) AddBook(c *gin.Context) {

	var book entities.Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 400,
			Message:    "the request body is not valid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := handler.Usecase.AddBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 500,
			Message:    "couldn't add new book",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 200,
		Message:    "successfully added the book",
		Data:       nil,
		Errors:     nil,
	})
}

func (handler *Handler) DeleteBook(c *gin.Context) {

	bookName := c.Param("bookName")

	if err := handler.Usecase.DeleteBook(bookName); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 500,
			Message:    "couldn't delete the book",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 200,
		Message:    "successfully deleted the book",
		Data:       nil,
		Errors:     nil,
	})
}
