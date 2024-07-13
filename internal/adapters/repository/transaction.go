package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
)

type TxRepository struct {
	MasterDB *sqlx.DB
}

func NewTxRepository(opts RepositoryOpts) ports.ITxRepository {
	return &TxRepository{
		MasterDB: opts.Database.MasterDB,
	}
}

// StartTx starts a new transaction.
func (m *TxRepository) StartTx(ctx context.Context) (*sql.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.StartTx")
	defer span.Finish()

	tx, err := m.MasterDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommitTx commits the given transaction.
func (m *TxRepository) CommitTx(ctx context.Context, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.CommitTx")
	defer span.Finish()

	return tx.Commit()
}

// RollbackTx rolls back the given transaction.
func (m *TxRepository) RollbackTx(ctx context.Context, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.RollbackTx")
	defer span.Finish()

	return tx.Rollback()
}
