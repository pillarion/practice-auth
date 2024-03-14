package auth

import (
	"context"

	"github.com/pillarion/practice-auth/internal/core/model/auth"
)

type Service interface {
	Login(ctx context.Context, cred auth.Credential) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
