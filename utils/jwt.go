package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const KEY = "supersecretKeyTest123213213"

func VerifyToken(token string) (int64, error) {
	if token == "" {
		return 0, errors.New("empty string token")
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(KEY), nil
	})
	if err != nil {
		return 0, err
	}

	validToken := parsedToken.Valid
	if !validToken {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}

func generateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(KEY))
}
