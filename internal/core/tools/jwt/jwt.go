package jwt

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	authModel "github.com/pillarion/practice-auth/internal/core/model/auth"
)

const (
	JWTAccessTokenVariant  = "access"
	JWTRefreshTokenVariant = "refresh"
)

// GenerateJWT generates a JWT token for the given user with the specified secret, duration, and token variant.
//
// Parameters:
// - user userModel.Info: the user information used to generate the JWT.
// - secret []byte: the secret key used to sign the JWT.
// - duration time.Duration: the duration for which the JWT will be valid.
// - tokenVariant string: the variant of the token to be generated.
// Returns:
// - string: the generated JWT token.
// - error: an error, if any, during token generation.
func GenerateJWT(username string, role string, secret []byte, duration time.Duration, tokenVariant string) (string, error) {
	claims := &authModel.Claims{
		Name:         username,
		Role:         role,
		TokenVariant: tokenVariant,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "practice-auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

// VerifyJWT verifies the JWT token with the provided secret and token variant.
//
// Parameters:
// - tokenString: a string representing the JWT token
// - secret: a byte array representing the secret key
// - tokenVariant: a string representing the token variant
// Returns:
// - *authModel.Claims: a pointer to the claims extracted from the JWT token
// - error: an error indicating any issues during verification
func VerifyJWT(tokenString string, secret []byte, tokenVariant string) (*authModel.Claims, error) {
	opt := []jwt.ParserOption{
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer("practice-auth"),
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&authModel.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
		opt...)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		if claims, ok := token.Claims.(*authModel.Claims); ok {
			if claims.TokenVariant == tokenVariant {
				return claims, nil
			}
		}
	}

	return nil, fmt.Errorf("invalid token")
}
