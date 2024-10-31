package util

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"noctiket/model/entity"
	"strings"
	"time"
)

type JwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

const Secret = "1n14d4l4hs3cr3t0k?"
const appName = "nocticket"

func GenerateJWT(user entity.User) string {
	secretKey := []byte(Secret)
	claims := JwtClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error signing token:", err)
		return ""
	}

	return tokenString
}

func GetClaims(authorizationToken string) (JwtClaims, error) {
	if authorizationToken == "" {
		return JwtClaims{}, nil
	}
	authorizationToken = strings.Replace(authorizationToken, "Bearer ", "", -1)
	var claims JwtClaims
	secretKey := []byte(Secret)

	_, err := jwt.ParseWithClaims(authorizationToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println(err.Error())
		return JwtClaims{}, err
	}

	return claims, nil
}
