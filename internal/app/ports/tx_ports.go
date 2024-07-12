package ports

import (
	"context"
	"database/sql"
)

//go:generate mockgen -source=./internal/app/ports/tx_ports.go -destination=./internal/mocks/mocks_tx.go -package=mocks

// IRepositoryTx defines the transaction methods.
type ITxRepository interface {
	StartTransaction(ctx context.Context) (*sql.Tx, error)
	CommitTransaction(ctx context.Context, tx *sql.Tx) error
	RollbackTransaction(ctx context.Context, tx *sql.Tx) error
}
