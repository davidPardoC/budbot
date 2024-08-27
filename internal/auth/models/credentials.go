package models

import "github.com/golang-jwt/jwt"

type Credentials struct {
	Token        string
	RefreshToken string
}

type TokenCustomClaims struct {
	ChatID   int64  `json:"chat_id"`
	UserId   int64  `json:"user_id"`
	PhotoUrl string `json:"photo_url"`
	jwt.StandardClaims
}
