package auth

import (
	"context"

	jwtTool "github.com/pillarion/practice-auth/internal/core/tools/jwt"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwtTool.VerifyJWT(refreshToken, []byte(s.jwtConfig.Secret), jwtTool.JWTRefreshTokenVariant)
	if err != nil {
		return "", err
	}

	accessToken, err := jwtTool.GenerateJWT(
		claims.Name,
		claims.Role,
		[]byte(s.jwtConfig.Secret),
		s.jwtConfig.AccessDuration,
		jwtTool.JWTAccessTokenVariant,
	)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
