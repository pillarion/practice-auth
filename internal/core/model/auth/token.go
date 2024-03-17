package auth

import "github.com/golang-jwt/jwt/v5"

type Credential struct {
	Username string
	Password string
}

type Claims struct {
	Name         string
	Role         string
	TokenVariant string
	jwt.RegisteredClaims
}
