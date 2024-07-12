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

// StartTransaction starts a new transaction.
func (m *TxRepository) StartTransaction(ctx context.Context) (*sql.Tx, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.StartTransaction")
	defer span.Finish()

	tx, err := m.MasterDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommitTransaction commits the given transaction.
func (m *TxRepository) CommitTransaction(ctx context.Context, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.CommitTransaction")
	defer span.Finish()

	return tx.Commit()
}

// RollbackTransaction rolls back the given transaction.
func (m *TxRepository) RollbackTransaction(ctx context.Context, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.TxRepository.RollbackTransaction")
	defer span.Finish()

	return tx.Rollback()
}
