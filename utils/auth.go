package utils

import (
	"context"
	"fmt"
	"greaterAltitudeapp/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))
var ctx = context.Background()

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"jti":    uuid.NewString(),
		"role":   role,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
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

	jti, ok := claims["jti"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'jti' claim")
	}

	userId, ok := claims["userId"].(uint)
        if !ok {
                return nil, fmt.Errorf("missing or invalid 'userId' claim")
        }

	if !ActiveUser(userId) {
		return nil, fmt.Errorf("Account is not active")
	}

	if IsJWTBlacklisted(jti) {
		return nil, fmt.Errorf("Token has been revoked")
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

func InvalidateJWT(tokenString string) error {

	claims, err := ValidateJWT(tokenString)

	if err != nil {
		return err
	}

	key := "Blacklisted:" + (*claims)["jti"].(string)
	exp := time.Unix(int64((*claims)["exp"].(float64)), 0)
	ttl := time.Until(exp)

	return H.RDB.Set(ctx, key, true, ttl).Err()
}

func IsJWTBlacklisted(checkJTI string) bool {
	key := "Blacklisted:" + checkJTI

	val, err := H.RDB.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	return val == "true"
}

func ActiveUser(userId uint) bool {
        var user models.User

        result := H.DB.First(&user, userId)
        if result.Error != nil {
            return false
        }

        return user.IsActive == true
}
