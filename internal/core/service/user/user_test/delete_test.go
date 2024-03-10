package user_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	minimock "github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/pillarion/practice-auth/internal/core/model/journal"
	journalRepo "github.com/pillarion/practice-auth/internal/core/port/repository/journal"
	journalRepoMock "github.com/pillarion/practice-auth/internal/core/port/repository/journal/mock"
	userRepo "github.com/pillarion/practice-auth/internal/core/port/repository/user"
	userRepoMock "github.com/pillarion/practice-auth/internal/core/port/repository/user/mock"
	service "github.com/pillarion/practice-auth/internal/core/service/user"
	txmanager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
	txmanagerMock "github.com/pillarion/practice-platform/pkg/pgtxmanager/mock"
	"github.com/stretchr/testify/require"
)

func Test_service_Delete(t *testing.T) {
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

		id   = int64(gofakeit.Number(1, 1000))
		jrnl = &journal.Journal{
			Action: "User deleted",
		}
		jid = int64(gofakeit.Number(1, 1000))
	)

	tests := []struct {
		name            string
		userRepoMock    userRepoMockFunc
		journalRepoMock journalRepoMockFunc
		transactor      txmFunc
		args            args
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func(mc *minimock.Controller) userRepo.Repo {
				mock := userRepoMock.NewRepoMock(mc).
					DeleteMock.Expect(ctxWithTx, id).
					Return(nil)

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
					DeleteMock.Expect(ctxWithTx, id).
					Return(nil)

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

			err := srvc.Delete(tt.args.ctx, tt.args.id)

			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
