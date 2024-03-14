package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	target "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	userService "github.com/pillarion/practice-auth/internal/core/port/service/user"
	serviceMock "github.com/pillarion/practice-auth/internal/core/port/service/user/mock"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestServer_Delete(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) userService.Service
	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}
	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		newID = int64(gofakeit.Number(1, 1000))
	)
	tests := []struct {
		name            string
		userServiceMock userServiceMockFunc
		args            args
		want            *emptypb.Empty
		wantErr         bool
	}{
		{
			name: "correctly delete user",
			userServiceMock: func(mc *minimock.Controller) userService.Service {
				m := serviceMock.NewServiceMock(mc).
					DeleteMock.Expect(ctx, newID).
					Return(nil)
				return m
			},
			args: args{
				ctx: ctx,
				req: &desc.DeleteRequest{
					Id: newID,
				},
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

			res, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, res)
		})
	}
}
