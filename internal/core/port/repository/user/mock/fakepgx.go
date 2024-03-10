package mock

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Pgtx struct{}

// Conn implements pgx.Tx.
func (p *Pgtx) Conn() *pgx.Conn {
	panic("unimplemented")
}

// CopyFrom implements pgx.Tx.
func (p *Pgtx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	panic("unimplemented")
}

// Exec implements pgx.Tx.
func (p *Pgtx) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	panic("unimplemented")
}

// LargeObjects implements pgx.Tx.
func (p *Pgtx) LargeObjects() pgx.LargeObjects {
	panic("unimplemented")
}

// Prepare implements pgx.Tx.
func (p *Pgtx) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	panic("unimplemented")
}

// Query implements pgx.Tx.
func (p *Pgtx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	panic("unimplemented")
}

// QueryRow implements pgx.Tx.
func (p *Pgtx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	panic("unimplemented")
}

// SendBatch implements pgx.Tx.
func (p *Pgtx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	panic("unimplemented")
}

func (*Pgtx) Commit(context.Context) error {
	return nil
}

func (*Pgtx) Rollback(context.Context) error {
	return nil
}

func (*Pgtx) Begin(context.Context) (pgx.Tx, error) {
	return &Pgtx{}, nil
}
