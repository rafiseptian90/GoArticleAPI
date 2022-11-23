package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Credentials struct {
	Email string
	*jwt.RegisteredClaims
}

func JWTGenerateToken(email string) (string, error) {
	expTime, _ := strconv.Atoi(os.Getenv("EXP_TIME"))
	expDuration := time.Now().Add(time.Duration(expTime) * time.Minute)

	claims := Credentials{
		Email:            email,
		RegisteredClaims: &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expDuration)},
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	token, err := jwtClaims.SignedString(jwtKey)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}

func JWTValidateToken(signedToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Credentials{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*Credentials); ok && token.Valid {
		return claims.Email, nil
	}

	return nil, errors.New(err.Error())
}

func JWTRefreshToken(email string, signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Credentials{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if _, ok := token.Claims.(*Credentials); ok && token.Valid {
		newToken, err := JWTGenerateToken(email)
		if err != nil {
			return "", errors.New(err.Error())
		}

		return newToken, nil
	} else {
		return "", errors.New(err.Error())
	}
}
