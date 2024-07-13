package ports

import (
	"context"
	"database/sql"
)

//go:generate mockgen -source=./internal/adapters/ports/tx_ports.go -destination=./internal/mocks/mocks_tx.go -package=mocks

// IRepositoryTx defines the transaction methods.
type ITxRepository interface {
	StartTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(ctx context.Context, tx *sql.Tx) error
	RollbackTx(ctx context.Context, tx *sql.Tx) error
}
