package database

import (
	"context"
	"fmt"

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

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Name)

	return &DatabaseOpts{
		MasterDB: connectMySQL(ctx, dsn, cfg.MaxOpenConns, cfg.MaxIdleConns),
	}
}

func connectMySQL(ctx context.Context, dsn string, maxOpenConns, maxIdleConns int) *sqlx.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "database.connectMySQL")
	defer span.Finish()

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.LogStdErr.Errorf("Failed to connect to MySQL: %s", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	return db
}
