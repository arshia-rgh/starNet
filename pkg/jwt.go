package pkg

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var ErrInvalidToken = errors.New("invalid token")

func GenerateToken(userID string, userRole string, secretKey string) (string, error) {
	if secretKey == "" {
		return "", fmt.Errorf("secret key can not be empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": userRole,
		"sub":  userID,
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string, secretKey string) (string, string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", "", err
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", "", ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid claims")
	}

	userID := claims["sub"].(string)
	userRole := claims["role"].(string)
	return userID, userRole, nil
}
