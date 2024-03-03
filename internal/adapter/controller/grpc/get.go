package grpc

import (
	"context"

	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Get implements the GetUser method of the UserV1Server interface.
func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	response := &desc.GetResponse{
		Id:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		Role:      desc.Role(desc.Role_value[res.Role]),
		CreatedAt: timestamppb.New(res.CreatedAt),
	}
	if !res.UpdatedAt.IsZero() {
		response.UpdatedAt = timestamppb.New(res.UpdatedAt)
	}

	return response, nil
}
