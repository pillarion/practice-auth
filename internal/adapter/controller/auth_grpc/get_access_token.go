package auth

import (
	"context"

	desc "github.com/pillarion/practice-auth/pkg/auth_v1"
)

// GetAccessToken implements the GetAccessToken method of the AuthV1Server interface.
func (s *server) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := s.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
