package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

const (
	queryGetUserByUsernameAndPassword = `SELECT id, name, username, email, password_hashed, is_active, created_at, updated_at 
		FROM users WHERE username=? AND password_hashed=?`
	queryGetEmailSame   = `SELECT username FROM users WHERE email=? `
	queryInsertDataUser = `INSERT INTO users (name,username,email,password_hashed,is_active,role_id) values (?,?,?,?,?,?)`
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

func (r *UserRepository) RegisterUser(ctx context.Context, userData entity.User) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.UserRepository.RegisterUser")
	defer span.Finish()

	user := entity.UserORM{}

	err := r.MasterDB.Get(&user, queryGetEmailSame, userData.Email)
	if err == nil {
		return -1, errors.New("Email Already Exists")
	}
	result, err := r.MasterDB.ExecContext(ctx, queryInsertDataUser, userData.Name, userData.Username, userData.Email, userData.Password, 1, 1)
	if err != nil {
		if result == nil {
			logger.LogStdErr.WithFields(logrus.Fields{
				"email": userData.Email,
				"error": err,
			}).Error("[UserRepository(RegisterUser)] " + err.Error())
			return -1, err
		}

	}
	lastid, err := result.LastInsertId()
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"email": userData.Email,
			"error": err,
		}).Error("[UserRepository(RegisterUser)] " + err.Error())
		return -1, err
	}
	return lastid, nil

}
