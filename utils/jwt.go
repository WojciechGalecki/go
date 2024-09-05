package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "changeit"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":   email,
			"user_id": userId,
			"exp":     time.Now().Add(time.Hour * 2).Unix(),
		},
	)

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

    if err != nil {
        return 0, errors.New("couldn't parse token")
    }

    isValid := parsedToken.Valid

    if !isValid {
        return 0, errors.New("invalid token")
    }

    claims, ok := parsedToken.Claims.(jwt.MapClaims)

    if !ok {
        return 0, errors.New("invalid token claims")
    }

    fmt.Println("---------")
    fmt.Println(claims)

    userId := int64(claims["user_id"].(float64))

    return userId, nil
}
