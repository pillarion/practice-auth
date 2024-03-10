package user_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"

	"github.com/pillarion/practice-auth/internal/core/model/journal"
	model "github.com/pillarion/practice-auth/internal/core/model/user"
	journalRepo "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	journalRepoMock "github.com/pillarion/practice-auth/internal/core/port/repository/journal/mock"
	userRepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	userRepoMock "github.com/pillarion/practice-auth/internal/core/port/repository/user/mock"
	service "github.com/pillarion/practice-auth/internal/core/service/user"
	txmanager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
	txmanagerMock "github.com/pillarion/practice-platform/pkg/pgtxmanager/mock"
	"github.com/stretchr/testify/require"
)

func Test_service_Get(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) userRepo.Repo
	type journalRepoMockFunc func(mc *minimock.Controller) journalRepo.Repo
	type txmFunc func(mc *minimock.Controller) txmanager.Transactor

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()

		txOption  = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		tx        = &userRepoMock.Pgtx{}
		ctxWithTx = context.WithValue(ctx, txmanager.TxKey, tx)

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
		jrnl = &journal.Journal{
			Action: "User get",
		}
		jid = int64(gofakeit.Number(1, 1000))
	)

	tests := []struct {
		name            string
		userRepoMock    userRepoMockFunc
		journalRepoMock journalRepoMockFunc
		transactor      txmFunc
		args            args
		want            *model.User
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name: "succes case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := userRepoMock.NewRepoMock(mc).
					SelectMock.Expect(ctxWithTx, id).
					Return(usr, nil)

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, jrnl).
					Return(jid, nil)

				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			want:    usr,
			wantErr: false,
		},
		{
			name: "failed case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := userRepoMock.NewRepoMock(mc).
					SelectMock.Expect(ctxWithTx, id).
					Return(usr, nil)

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, jrnl).
					Return(0, gofakeit.Error())

				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed case v2",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := userRepoMock.NewRepoMock(mc).
					SelectMock.Expect(ctxWithTx, id).
					Return(nil, gofakeit.Error())

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc)

				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := minimock.NewController(t)
			userRepo := tt.userRepoMock(mc)
			journalRepo := tt.journalRepoMock(mc)
			txManager := txmanager.NewTransactionManager(tt.transactor(mc))

			srvc := service.NewService(userRepo, journalRepo, txManager)

			res, err := srvc.Get(ctx, tt.args.id)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.wantErr, err != nil)

		})
	}
}
