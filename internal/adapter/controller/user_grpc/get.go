package grpc

import (
	"context"

	dto "github.com/pillarion/practice-auth/internal/core/dto/grpc"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

// Get implements the GetUser method of the UserV1Server interface.
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	d := dto.UserToDTO(res)

	response := &desc.GetResponse{
		Id:        d.User.ID,
		Name:      d.User.Name,
		Email:     d.User.Email,
		Role:      d.User.Role,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	return response, nil
}
