package user_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	userRepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	repoMock "github.com/pillarion/practice-auth/internal/core/port/repository/user/mock"
	"github.com/stretchr/testify/require"
)

func Test_service_Delete(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) userRepo.Repo
	type args struct {
		ctx context.Context
		id  int64
	}
	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()
		id  = int64(gofakeit.Number(1, 1000))
	)
	tests := []struct {
		name         string
		userRepoMock userRepoMockFunc
		args         args
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := repoMock.NewRepoMock(mc).
					DeleteMock.Expect(ctx, id).
					Return(nil)

				return mock
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := tt.userRepoMock(mc)
			err := repo.Delete(tt.args.ctx, tt.args.id)

			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
