package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	target "github.com/pillarion/practice-auth/internal/adapter/controller/user_grpc"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
	userService "github.com/pillarion/practice-auth/internal/core/port/service/user"
	serviceMock "github.com/pillarion/practice-auth/internal/core/port/service/user/mock"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestServer_Update(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) userService.Service
	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}
	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name  = gofakeit.Name()
		email = gofakeit.Email()

		newID = int64(gofakeit.Number(1, 1000))

		req = &desc.UpdateRequest{
			Id:    newID,
			Name:  wrapperspb.String(name),
			Email: wrapperspb.String(email),

			Role: desc.Role(desc.Role_value["ADMIN"]),
		}

		usr = &model.Info{
			ID:    newID,
			Name:  name,
			Email: email,
			Role:  "ADMIN",
		}
	)
	tests := []struct {
		name            string
		userServiceMock userServiceMockFunc
		args            args
		want            *emptypb.Empty
		wantErr         bool
	}{
		{
			name: "correctly update user",
			userServiceMock: func(mc *minimock.Controller) userService.Service {
				m := serviceMock.NewServiceMock(mc).
					UpdateMock.Expect(ctx, usr).
					Return(nil)
				return m
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := target.NewServer(tt.userServiceMock(mc))

			res, err := api.Update(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, res)
		})
	}
}
