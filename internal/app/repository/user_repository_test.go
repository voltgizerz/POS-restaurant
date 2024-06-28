package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

func TestUserRepository_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name     string
		r        *UserRepository
		args     args
		want     *entity.UserORM
		wantErr  bool
		mockFunc func(mock sqlmock.Sqlmock)
	}{
		{
			name: "SUCCESS - GetUserByEmail",
			args: args{

				ctx:   context.Background(),
				email: "mock@email.com",
			},
			want: &entity.UserORM{
				ID: 1,
			},
			wantErr: false,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(queryGetEmailSame).
					WithArgs("mock@email.com").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
		},
		{
			name: "ERROR - On GetUserByEmail",
			args: args{
				ctx:   context.Background(),
				email: "mock@email.com",
			},
			want:    nil,
			wantErr: true,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(queryGetEmailSame).
					WithArgs("mock@email.com").
					WillReturnError(errors.New("some error"))
			},
		},
		{
			name: "ERROR - On GetUserByEmail got NoRows",
			args: args{
				ctx:   context.Background(),
				email: "mock@email.com",
			},
			want:    &entity.UserORM{},
			wantErr: false,
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(queryGetEmailSame).
					WithArgs("mock@email.com").
					WillReturnError(sql.ErrNoRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error creating mock database: %v", err)
			}
			defer dbMock.Close()

			// Initialize UserRepository with the mock DB
			tt.r = &UserRepository{
				MasterDB: sqlx.NewDb(dbMock, "sqlmock"),
			}

			// Call the mock function to set expectations
			tt.mockFunc(mock)

			// Call the method under test
			got, err := tt.r.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByEmail() = %v, want %v", got, tt.want)
			}

			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_GetUserByUsernameAndPassword(t *testing.T) {
	type args struct {
		ctx          context.Context
		username     string
		hashPassword string
	}
	tests := []struct {
		name    string
		r       *UserRepository
		args    args
		want    *entity.UserORM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUserByUsernameAndPassword(tt.args.ctx, tt.args.username, tt.args.hashPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByUsernameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByUsernameAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
