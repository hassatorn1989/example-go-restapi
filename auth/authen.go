package auth

import (
	"inv-app/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID string, isRefresh bool) (string, error) {
	if isRefresh {
		claims := jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Minute * 15).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(configs.JWT_SECRET_ACCESS))
	} else {
		claims := jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Minute * 15).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(configs.JWT_SECRET_REFRESH))
	}
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(configs.JWT_SECRET_ACCESS), nil
	})
}

func RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(configs.JWT_SECRET_ACCESS), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		return token.SignedString([]byte(configs.JWT_SECRET_ACCESS))
	}
	return "", jwt.ErrSignatureInvalid
}
