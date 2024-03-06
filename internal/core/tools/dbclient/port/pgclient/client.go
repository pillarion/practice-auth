package port

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	pgtxmanager "github.com/pillarion/practice-auth/internal/core/tools/dbclient/port/pgtxmanager"
)

// Client represents a client for DB
type Client interface {
	DB() DB
	Close() error
}

// DB represents a database
type DB interface {
	SQLExecer
	pgtxmanager.Transactor
	Pinger
	Close()
}

// Pinger represents a pinger
type Pinger interface {
	Ping(ctx context.Context) error
}

// SQLExecer represents an executor
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer represents a named query scanner
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer represents a query executor
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Query represents a query wrapper
type Query struct {
	Name     string
	QueryRaw string
}
