package usermodel

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
