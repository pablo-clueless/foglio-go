package common

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var (
	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET"))
	JwtSigningMethod = jwt.SigningMethodHS256.Name
)

func JwtGenerateToken(Id string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        Id,
		Issuer:    Id,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(JwtSecretKey)
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return JwtSecretKey, nil
}

func JwtVerifyToken(str string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(str, claims, validateSignedMethod)
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid auth token")
	}

	return claims, nil
}
