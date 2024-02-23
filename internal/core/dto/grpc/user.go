package grpc

import (
	pb "github.com/pillarion/practice-auth/pkg/user_v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// UserInfoDTO defines the user model for the grpc server
type UserInfoDTO struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Role     pb.Role
}

// UserDTO defines the user model for the grpc server
type UserDTO struct {
	User      *UserInfoDTO
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}
