package helpers

import (
	res "api/types/structs/responses"
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