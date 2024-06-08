package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type UserRepository struct {
	MasterDB *sqlx.DB
}

func NewUserRepository(opts RepositoryOpts) ports.IUserRepository {
	return &UserRepository{
		MasterDB: opts.Database.MasterDB,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, userID int64) (*entity.User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "repo.UserRepository.GetUser")
	defer span.Finish()

	user := entity.User{}
	err := r.MasterDB.Get(&user, "SELECT * FROM users WHERE id=$1", userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[UserRepository] error on GetUser")

		return nil, err
	}

	return &user, nil
}
