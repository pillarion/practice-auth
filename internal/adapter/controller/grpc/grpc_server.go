package grpc

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedUserV1Server
	userService user.Service
}

// NewServer creates a new server instance.
//
// No parameters. Returns a pointer to a server.
func NewServer(us user.Service) *server {
	return &server{
		userService: us,
	}
}

// UpdateUser implements the UpdateUser method of the UserV1Server interface.
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	updUser := &model.User{
		ID:    req.GetId(),
		Name:  req.GetName().GetValue(),
		Email: req.GetEmail().GetValue(),
		Role:  req.GetRole().String(),
	}

	err := s.userService.Update(ctx, updUser)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteUser implements the DeleteUser method of the UserV1Server interface.
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := s.userService.Delete(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
