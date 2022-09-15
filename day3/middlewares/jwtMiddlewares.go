package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type MyCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ValidateToken(encodedToken string) (int, error) {
	mySigningKey := []byte(os.Getenv("SECRET_JWT"))
	token, err := jwt.ParseWithClaims(encodedToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("token invalid")
	}

}
