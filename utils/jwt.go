package utils

import (
	"GoPay/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("4QATynKRjKnMlZp9jy28KO9BTG8gjBBfSS4TjlpUtb4=")

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	user.Token = tokenString

	return tokenString, nil
}
