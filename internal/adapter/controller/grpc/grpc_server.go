package grpc

import (
	"context"

	"github.com/brianvoe/gofakeit"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
	"github.com/pillarion/practice-auth/internal/core/port/service/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// implementation of UserV1Server
// GetUser implements the GetUser method of the UserV1Server interface.
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	res, err := s.userService.Get(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Id:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		Role:      desc.Role(res.Role),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// CreateUser implements the CreateUser method of the UserV1Server interface.
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	user := &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     int(req.GetRole()),
	}

	res, err := s.userService.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: res,
	}, nil
}

// UpdateUser implements the UpdateUser method of the UserV1Server interface.
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	updUser := &model.User{
		ID:    req.GetId(),
		Name:  req.GetName().GetValue(),
		Email: req.GetEmail().GetValue(),
		Role:  int(req.GetRole()),
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
