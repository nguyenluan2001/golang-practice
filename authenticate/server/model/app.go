package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type JWTToken struct {
	Email     string
	CreatedAt string
	jwt.RegisteredClaims
}
