package utils

import (
	"fmt"
	"github.com/BerkatPS/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID int) (string, error) {
	cfg := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(cfg.JwtSecret))
}

func ValidateToken(tokenString string) bool {
	cfg := config.LoadConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false
	}
	return true
}
