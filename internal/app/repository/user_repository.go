package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

const (
	queryGetUserByUsernameAndPassword = `SELECT id, name, username, email, password_hashed, is_active, created_at, updated_at 
		FROM users WHERE username=? AND password_hashed=?`
	queryGetEmailSame   = `SELECT username FROM users WHERE email=? `
	queryInsertDataUser = `INSERT INTO (name,username,email,password_hashed,is_active,created_at) values (?,?,?,?,?,?)`
)

type UserRepository struct {
	MasterDB *sqlx.DB
}

func NewUserRepository(opts RepositoryOpts) ports.IUserRepository {
	return &UserRepository{
		MasterDB: opts.Database.MasterDB,
	}
}

func (r *UserRepository) GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*entity.UserORM, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.UserRepository.GetUserByUsernameAndPassword")
	defer span.Finish()

	user := entity.UserORM{}
	err := r.MasterDB.GetContext(ctx, &user, queryGetUserByUsernameAndPassword, username, hashPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) RegisterUser(ctx context.Context, username string, hashPassword string, Email string, Name string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.UserRepository.RegisterUser")
	defer span.Finish()

	user := entity.UserORM{}

	err := r.MasterDB.Get(&user, queryGetEmailSame, Email)
	if err == nil {
		return false, err
	}

	result := r.MasterDB.MustExecContext(ctx, queryInsertDataUser, Name, username, Email, hashPassword, 1, time.Now())
	if result == nil {
		return false, errors.New("Failed Insert User On Database")
	}

	return true, nil

}
