package grpc

import (
	pb "github.com/pillarion/practice-auth/pkg/user_v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type UserInfoDTO struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Role     pb.Role
}

type UserDTO struct {
	User      *UserInfoDTO
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}
