package middlewares

import (
	"errors"
	"net/http"

	customerrormessages "github.com/akshay0074700747/book_store/custom_errormessages"
	jwttoken "github.com/akshay0074700747/book_store/web/jwt_token"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	secret string
}

func NewMiddleware(secret string) *Middleware {
	return &Middleware{
		secret: secret,
	}
}

func (middleware *Middleware) GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("Token")
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Cookie cannot be retrieved")
			c.Abort()
			return
		}

		if cookie == nil {
			c.JSON(http.StatusUnauthorized, "Please login")
			c.Abort()
			return
		}

		values, err := jwttoken.ValidateToken(cookie.Value, []byte(middleware.secret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("values", values)
	}
}

func (middleware *Middleware) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		value, exists := c.Get("values")
		if !exists {
			c.AbortWithError(http.StatusInternalServerError, errors.New(customerrormessages.CredentialsNotfound))
			return
		}

		valueMap, _ := value.(map[string]interface{})

		isAdmin := valueMap["isAdmin"].(bool)

		if !isAdmin {
			c.JSON(http.StatusUnauthorized, "this route is only accessible to admins")
			c.Abort()
			return
		}

	}
}
