package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var ErrInvalidToken = errors.New("invalid token")

func GenerateToken(userID int, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string, secretKey string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	userID := int(claims["sub"].(float64))
	return userID, nil
}
