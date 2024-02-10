package grpc

import (
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

type server struct {
	desc.UnimplementedUserV1Server
}

// NewServer creates a new server instance.
//
// No parameters. Returns a pointer to a server.
func NewServer() *server {
	return &server{
		UnimplementedUserV1Server: desc.UnimplementedUserV1Server{},
	}
}
