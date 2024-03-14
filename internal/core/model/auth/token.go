package auth

import "github.com/golang-jwt/jwt/v5"

type Credential struct {
	Username string
	Password string
}

type Claims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}
