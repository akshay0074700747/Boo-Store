package handlers

import (
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

type Response struct {
	StatusCode int         `json:"stastuscode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"error,omitempty"`
}

func NewHandler(usecase usecaseports.UsecasePort, secret string) *Handler {

	return &Handler{
		secret:  secret,
		Usecase: usecase,
	}
}

func (handler *Handler) Login(c *gin.Context) {

	cookie, err := c.Request.Cookie("Token")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if cookie != nil {
		c.JSON(http.StatusConflict, Response{
			StatusCode: 409,
			Message:    "the user is already logged in",
			Data:       nil,
			Errors:     nil,
		})
		return
	}

	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 400,
			Message:    "cannot bind the request body",
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

	c.SetCookie("jwtToken", jwt, 3600, "/", "localhost", false, true)
	result.Password = ""

	c.JSON(http.StatusOK, Response{
		StatusCode: 200,
		Message:    "user logged in successfully",
		Data:       result,
		Errors:     nil,
	})
}
