package util

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "JWT-SecretKey"
)

func GenerateJWTToken(userID string) string {

	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken
}

func GetUserIDFromToken(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return "", fmt.Errorf("invalid token signing method")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user ID format")
	}

	return userID, nil
}
