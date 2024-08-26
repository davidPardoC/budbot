package models

import "github.com/golang-jwt/jwt"

type Credentials struct {
	Token        string
	RefreshToken string
}

type TokenCustomClaims struct {
	ChatID   int64
	UserId   int64
	PhotoUrl string
	jwt.StandardClaims
}
