package jwttoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type payload struct {
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

// generating jwt token from the given credentials
func GenerateJwt(userName string, isAdmin bool, secret []byte) (string, error) {
	expiresAt := time.Now().Add(48 * time.Hour)
	jwtClaims := &payload{
		UserName: userName,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}

// validating the token
func ValidateToken(tokenstring string, secret []byte) (map[string]interface{}, error) {

	token, err := jwt.ParseWithClaims(tokenstring, &payload{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid token")
		}

		return secret, nil

	})

	if err != nil {
		return nil, err
	}

	if token == nil || !token.Valid {
		return nil, fmt.Errorf("token is not valid or its empty")
	}

	cliams, ok := token.Claims.(*payload)

	if !ok {
		return nil, fmt.Errorf("cannot parse claims")
	}

	cred := map[string]interface{}{
		"userName": cliams.UserName,
		"isAdmin":  cliams.IsAdmin,
	}

	if cliams.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return cred, nil

}
