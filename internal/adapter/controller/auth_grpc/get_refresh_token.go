package auth_grpc

import (
	"context"

	desc "github.com/pillarion/practice-auth/pkg/auth_v1"
)

// Get oldRefreshToken implements the GetoldRefreshToken method of the AuthV1Server interface.
func (s *server) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	newRefreshToken, err := s.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: newRefreshToken,
	}, nil
}
