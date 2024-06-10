package service

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/internal/utils"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type UserService struct {
	authService    ports.IAuth
	userRepository ports.IUserRepository
}

func NewUserService(i interactor.UserService) *UserService {
	return &UserService{
		authService:    i.AuthService,
		userRepository: i.UserRepository,
	}
}

func (s *UserService) Login(ctx context.Context, username string, password string) (*entity.LoginResponse, error) {
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

	tokenData, err := s.authService.CreateToken(user)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on CreateToken")

		return nil, err
	}

	resp := &entity.LoginResponse{
		UserID:    user.ID,
		RoleID:    1, // TODO
		Token:     tokenData.Token,
		TokenType: tokenData.TokenType,
		ExpiredAt: tokenData.ExpiredAt,
	}

	return resp, nil
}

func (s *UserService) Register(ctx context.Context, userData entity.UserORM) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UserService.Register")
	defer span.Finish()

	passwordhasing, err := utils.HashPassword(userData.Password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on UserService Hash Password On Register")
	}
	userData.Password = passwordhasing

	result, err := s.userRepository.RegisterUser(ctx, userData)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on UserService From Repository Register User")
		return 0, err
	}

	return result, nil
}
