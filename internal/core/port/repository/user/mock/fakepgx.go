package mock

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Pgtx implements pgx.Tx.
type Pgtx struct{}

// Conn implements pgx.Tx.
func (p *Pgtx) Conn() *pgx.Conn {
	panic("unimplemented")
}

// CopyFrom implements pgx.Tx.
func (p *Pgtx) CopyFrom(_ context.Context, _ pgx.Identifier, _ []string, _ pgx.CopyFromSource) (int64, error) {
	panic("unimplemented")
}

// Exec implements pgx.Tx.
func (p *Pgtx) Exec(_ context.Context, _ string, _ ...any) (commandTag pgconn.CommandTag, err error) {
	panic("unimplemented")
}

// LargeObjects implements pgx.Tx.
func (p *Pgtx) LargeObjects() pgx.LargeObjects {
	panic("unimplemented")
}

// Prepare implements pgx.Tx.
func (p *Pgtx) Prepare(_ context.Context, _ string, _ string) (*pgconn.StatementDescription, error) {
	panic("unimplemented")
}

// Query implements pgx.Tx.
func (p *Pgtx) Query(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
	panic("unimplemented")
}

// QueryRow implements pgx.Tx.
func (p *Pgtx) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	panic("unimplemented")
}

// SendBatch implements pgx.Tx.
func (p *Pgtx) SendBatch(_ context.Context, _ *pgx.Batch) pgx.BatchResults {
	panic("unimplemented")
}

// Commit implements pgx.Tx.
func (*Pgtx) Commit(context.Context) error {
	return nil
}

// Rollback implements pgx.Tx.
func (*Pgtx) Rollback(context.Context) error {
	return nil
}

// Begin implements pgx.Tx.
func (*Pgtx) Begin(context.Context) (pgx.Tx, error) {
	return &Pgtx{}, nil
}
