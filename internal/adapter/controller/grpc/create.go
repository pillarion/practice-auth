package grpc

import (
	"context"
	"fmt"

	model "github.com/pillarion/practice-auth/internal/core/model/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

// CreateUser implements the CreateUser method of the UserV1Server interface.
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.GetName() == "" || req.GetEmail() == "" || req.GetPassword() == "" || req.GetRole().Number() == 0 {

		return nil, fmt.Errorf("name, email, password and role are required")
	}

	if req.GetPassword() != req.GetPasswordConfirm() {

		return nil, fmt.Errorf("passwords do not match")
	}

	user := &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole().String(),
	}

	res, err := s.userService.Create(ctx, user)
	if err != nil {

		return nil, err
	}

	return &desc.CreateResponse{
		Id: res,
	}, nil
}
