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

func Test_service_Create(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) userRepo.Repo

	type args struct {
		ctx  context.Context
		user *model.Info
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 8)

		usr = &model.Info{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     "ADMIN",
		}

		newID = int64(gofakeit.Number(1, 1000))
	)
	tests := []struct {
		name         string
		userRepoMock userRepoMockFunc
		args         args
		want         int64
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := repoMock.NewRepoMock(mc).
					InsertMock.Expect(ctx, usr).
					Return(newID, nil)

				return mock
			},
			args: args{
				ctx:  ctx,
				user: usr,
			},
			want:    newID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := tt.userRepoMock(mc)
			newID, err := repo.Insert(tt.args.ctx, tt.args.user)

			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
