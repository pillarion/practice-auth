package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	target "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
	userService "github.com/pillarion/practice-auth/internal/core/port/service/user"
	serviceMock "github.com/pillarion/practice-auth/internal/core/port/service/user/mock"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer_Get(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) userService.Service
	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}
	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name      = gofakeit.Name()
		email     = gofakeit.Email()
		password  = gofakeit.Password(true, true, true, true, true, 8)
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
		ID        = int64(gofakeit.Number(1, 1000))

		usr = &model.User{
			Info: model.Info{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     "ADMIN",
				ID:       ID,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	)
	tests := []struct {
		name            string
		userServiceMock userServiceMockFunc
		args            args
		want            *desc.GetResponse
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name: "correctly get user",
			userServiceMock: func(mc *minimock.Controller) userService.Service {
				m := serviceMock.NewServiceMock(mc).
					GetMock.Expect(ctx, ID).
					Return(usr, nil)
				return m
			},
			args: args{
				ctx: ctx,
				req: &desc.GetRequest{
					Id: ID,
				},
			},
			want: &desc.GetResponse{
				Id:        usr.Info.ID,
				Name:      usr.Info.Name,
				Email:     usr.Info.Email,
				Role:      desc.Role(desc.Role_value[usr.Info.Role]),
				CreatedAt: timestamppb.New(usr.CreatedAt),
				UpdatedAt: timestamppb.New(usr.UpdatedAt),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := target.NewServer(tt.userServiceMock(mc))

			res, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, res)
		})
	}
}
