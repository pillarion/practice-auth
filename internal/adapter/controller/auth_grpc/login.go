package auth_grpc

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/auth"
	desc "github.com/pillarion/practice-auth/pkg/auth_v1"
)

// Login implements the Login method of the AuthV1Server interface.
func (s *server) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	creds := model.Credential{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	res, err := s.authService.Login(ctx, creds)
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: res,
	}, nil
}
