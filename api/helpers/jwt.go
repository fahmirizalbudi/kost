package helpers

import (
	res "api/types/structs/responses"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    	int		`json:"id"`
	Name  	string	`json:"name"`
	Email 	string	`json:"email"`
	Role	string	`json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user res.UserResponse) (string, error) {
	claims := Claims{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, errors.New("unauthorized")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	return claims, nil
}