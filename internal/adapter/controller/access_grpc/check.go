package access_grpc

import (
	"context"
	"fmt"

	desc "github.com/pillarion/practice-auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Check access
func (s *server) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	err := s.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, fmt.Errorf("access denied: %v", err)
	}

	return &emptypb.Empty{}, nil
}
