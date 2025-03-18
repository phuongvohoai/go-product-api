package services

import (
	"fmt"
	"os"
	"phuong/go-product-api/database"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, email string) (string, error) {
	expiredAt := time.Now().Add(time.Hour * 24).Unix()
	jwtId, _ := uuid.NewV7()
	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Id:        jwtId.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func VerifyToken(token string) (bool, Claims) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return false, Claims{}
	}

	return true, *claims
}

func RevokeToken(jti string) error {
	// Add user to black list
	blackListKey := fmt.Sprintf("BLACKLIST_TOKEN_%s", jti)
	err := database.Redis.Set(blackListKey, true, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func IsTokenRevoked(jti string) bool {
	blackListKey := fmt.Sprintf("BLACKLIST_TOKEN_%s", jti)
	_, err := database.Redis.Get(blackListKey).Result()
	if err != nil {
		return false
	}

	return true
}
