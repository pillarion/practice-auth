package user_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"

	model "github.com/pillarion/practice-auth/internal/core/model/user"
	userRepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	repoMock "github.com/pillarion/practice-auth/internal/core/port/repository/user/mock"
	"github.com/stretchr/testify/require"
)

func Test_service_Get(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) userRepo.Repo
	type args struct {
		ctx context.Context
		id  int64
	}
	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 8)
		id       = int64(gofakeit.Number(1, 1000))

		usr = &model.User{
			Info: model.Info{
				ID:       id,
				Name:     name,
				Email:    email,
				Password: password,
				Role:     "ADMIN",
			},
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		}
	)
	tests := []struct {
		name         string
		userRepoMock userRepoMockFunc
		args         args
		want         *model.User
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
					SelectMock.Expect(ctx, id).
					Return(usr, nil)

				return mock
			},
			want:    usr,
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := repoMock.NewRepoMock(mc).
					SelectMock.Expect(ctx, id).
					Return(nil, gofakeit.Error())

				return mock
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := tt.userRepoMock(mc)
			newID, err := repo.Select(tt.args.ctx, tt.args.id)

			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
