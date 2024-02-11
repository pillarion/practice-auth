package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	desc.UnimplementedUserV1Server
}

// NewServer creates a new server instance.
//
// No parameters. Returns a pointer to a server.
func NewServer() *server {
	return &server{}
}

// implementation of UserV1Server
// GetUser implements the GetUser method of the UserV1Server interface.
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	slog.Info("GetUser", "request", req.String())

	return &desc.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// CreateUser implements the CreateUser method of the UserV1Server interface.
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("CreateUser", "request", req.String())

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// UpdateUser implements the UpdateUser method of the UserV1Server interface.
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	slog.Info("UpdateUser", "request", req.String())

	return &emptypb.Empty{}, nil
}

// DeleteUser implements the DeleteUser method of the UserV1Server interface.
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("DeleteUser", "request", req.String())

	return &emptypb.Empty{}, nil
}
