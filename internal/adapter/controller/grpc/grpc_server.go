package grpc

import (
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

type server struct {
	desc.UnimplementedUserV1Server
	userService user.Service
}

func NewServer(us user.Service) *server {
	return &server{
		userService: us,
	}
}
