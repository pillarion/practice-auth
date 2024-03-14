package grpc

import (
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

// Server implements the user gRPC service.
type server struct {
	desc.UnimplementedUserV1Server
	userService user.Service
}

// NewServer creates a new server instance with the given user service.
// It takes a user service as a parameter and returns a pointer to a server.
func NewServer(us user.Service) *server {
	return &server{
		userService: us,
	}
}
