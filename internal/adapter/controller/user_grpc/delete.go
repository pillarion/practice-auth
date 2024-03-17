package user_grpc

import (
	"context"

	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete implements the DeleteUser method of the UserV1Server interface.
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
