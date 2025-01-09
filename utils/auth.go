package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtSecret = os.Getenv("SECRET_KEY")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return &claims, nil
}

func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashed), nil
}
