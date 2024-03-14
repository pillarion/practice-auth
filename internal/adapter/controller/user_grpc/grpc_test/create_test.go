package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	target "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
	userService "github.com/pillarion/practice-auth/internal/core/port/service/user"
	serviceMock "github.com/pillarion/practice-auth/internal/core/port/service/user/mock"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
)

func TestServer_Create(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) userService.Service

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 8)

		newID = int64(gofakeit.Number(1, 1000))

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role(desc.Role_value["ADMIN"]),
		}

		usr = &model.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     "ADMIN",
		}

		res = &desc.CreateResponse{
			Id: newID,
		}
	)

	tests := []struct {
		name            string
		userServiceMock userServiceMockFunc
		args            args
		want            *desc.CreateResponse
		wantErr         bool
	}{
		{
			name: "create user",
			userServiceMock: func(mc *minimock.Controller) userService.Service {
				mock := serviceMock.NewServiceMock(mc).
					CreateMock.Expect(ctx, usr).
					Return(newID, nil)

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    res,
			wantErr: false,
		},
		{
			name: "create user error",
			userServiceMock: func(mc *minimock.Controller) userService.Service {
				mock := serviceMock.NewServiceMock(mc).
					CreateMock.Expect(ctx, usr).
					Return(0, gofakeit.Error())

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := target.NewServer(tt.userServiceMock(mc))
			newID, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, newID)
		})
	}
}
