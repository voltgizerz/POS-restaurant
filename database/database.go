package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type DatabaseOpts struct {
	MasterDB *sqlx.DB
}

// InitDatabase initializes and returns a new Database instance
func InitDatabase(ctx context.Context, cfg config.Database) *DatabaseOpts {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database.InitDatabase")
	defer span.Finish()

	return &DatabaseOpts{
		MasterDB: connectMySQL(ctx, cfg),
	}
}

func connectMySQL(ctx context.Context, cfg config.Database) *sqlx.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "database.connectMySQL")
	defer span.Finish()

	db, err := sqlx.Connect("mysql", cfg.Host)
	if err != nil {
		logger.LogStdErr.Errorf("Failed to connect to MySQL: %s", err)
	}

	return db
}
