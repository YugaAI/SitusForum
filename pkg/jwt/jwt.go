package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(time.Minute * 10).Unix(),
		},
	)

	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claim["id"].(float64)), claim["username"].(string), nil
}

func ValidateTokenWithoutExpiry(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claim["id"].(float64)), claim["username"].(string), nil
}
