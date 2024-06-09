package service

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/internal/utils"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type UserService struct {
	userRepository ports.IUserRepository
}

func NewUserService(userRepository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UserService.Login")
	defer span.Finish()

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on HashPassword")

		return nil, err
	}

	user, err := s.userRepository.GetUserByUsernameAndPassword(ctx, username, hashPassword)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on GetUserByUsernameAndPassword")

		return nil, err
	}

	userData := &entity.User{
		ID: user.ID,
	}

	return userData, nil
}

func (s *UserService) Register(ctx context.Context, username string, email string, password string, confirmPass string, name string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UserService.Register")
	defer span.Finish()

	if password != confirmPass {
		return false, errors.New("Password Mismatch")
	}

	passwordhasing, err := utils.HashPassword(password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("[UserService] error on UserService Register")
	}

	result, err := s.userRepository.RegisterUser(ctx, username, passwordhasing, email, name)
	if !result {
		return false, err
	}

	return true, nil
}
