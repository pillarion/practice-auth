package grpc

import (
	"context"

	model "github.com/pillarion/practice-auth/internal/core/model/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update implements the UpdateUser method of the UserV1Server interface.
func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
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
