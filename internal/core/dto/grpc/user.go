package grpc

import (
	model "github.com/pillarion/practice-auth/internal/core/model/user"
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

// UserToDTO converts the user model to the grpc dto
func UserToDTO(user *model.User) *UserDTO {
	dto := &UserDTO{
		User: &UserInfoDTO{
			ID:       user.Info.ID,
			Name:     user.Info.Name,
			Email:    user.Info.Email,
			Password: user.Info.Password,
			Role:     pb.Role(pb.Role_value[user.Info.Role]),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	if !user.UpdatedAt.IsZero() {
		dto.UpdatedAt = timestamppb.New(user.UpdatedAt)
	}

	return dto
}

// UserToModel converts the grpc dto to the user model
func UserToModel(dto *UserDTO) *model.User {
	return &model.User{
		Info: model.Info{
			ID:       dto.User.ID,
			Name:     dto.User.Name,
			Email:    dto.User.Email,
			Password: dto.User.Password,
			Role:     dto.User.Role.String(),
		},
		CreatedAt: dto.CreatedAt.AsTime(),
		UpdatedAt: dto.UpdatedAt.AsTime(),
	}
}

// UserToModelInfo converts the grpc dto to the UserInfo model
func UserToModelInfo(dto *UserDTO) *model.Info {
	return &model.Info{
		ID:       dto.User.ID,
		Name:     dto.User.Name,
		Email:    dto.User.Email,
		Password: dto.User.Password,
		Role:     dto.User.Role.String(),
	}
}
