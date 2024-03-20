package auth

import "github.com/golang-jwt/jwt/v5"

// Credential holds the username and password.
type Credential struct {
	Username string
	Password string
}

// Claims holds the token claims.
type Claims struct {
	Name         string
	Role         string
	TokenVariant string
	jwt.RegisteredClaims
}
