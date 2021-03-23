package utils

import (
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(jwtToken string) (bool, jwt.MapClaims) {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if token.Valid && err == nil {
		return true, tokenData
	} else {
		return false, tokenData
	}
}