package auth

import (
	"github.com/pillarion/practice-auth/internal/core/port/service/auth"
	desc "github.com/pillarion/practice-auth/pkg/auth_v1"
)

type server struct {
	desc.UnimplementedAuthV1Server
	authService auth.Service
}

// NewServer creates a new server instance with the given auth service.
func NewServer(as auth.Service) *server {
	return &server{
		authService: as,
	}
}
