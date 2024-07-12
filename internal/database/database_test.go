package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Init()
	m.Run()
}
func TestInitDatabase(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg config.Database
	}
	tests := []struct {
		name string
		args args
		want *DatabaseOpts
	}{
		{
			name: "Successful database initialization",
			args: args{
				ctx: context.Background(),
				cfg: config.Database{
					Username:     "testuser",
					Password:     "testpass",
					Host:         "localhost:3306",
					Name:         "testdb",
					MaxOpenConns: 10,
					MaxIdleConns: 1,
				},
			},
			want: &DatabaseOpts{
				MasterDB: func() *sqlx.DB {
					db, _, _ := sqlmock.New() // create a new sqlmock database

					sqlxDB := sqlx.NewDb(db, "sqlmock")
					return sqlxDB
				}(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Override sqlxConnectContext with a mock implementation
			sqlxConnectContext = func(ctx context.Context, driverName, dataSourceName string) (*sqlx.DB, error) {
				db, _, err := sqlmock.New()
				if err != nil {
					t.Fatalf("failed to create sqlmock: %v", err)
				}
				return sqlx.NewDb(db, "sqlmock"), nil
			}

			got := InitDatabase(tt.args.ctx, tt.args.cfg)
			if !reflect.DeepEqual(got.MasterDB.DriverName(), tt.want.MasterDB.DriverName()) {
				t.Errorf("InitDatabase() = %v, want %v", got.MasterDB.DriverName(), tt.want.MasterDB.DriverName())
			}

			assert.Equal(t, tt.args.cfg.MaxOpenConns, got.MasterDB.Stats().MaxOpenConnections)
			assert.Equal(t, tt.args.cfg.MaxIdleConns, got.MasterDB.Stats().Idle)
		})
	}
}
