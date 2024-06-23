package service

import (
	"context"
	"errors"

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

	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on GetUserByUsername")

		return nil, err
	}

	err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on VerifyPassword")

		return nil, err
	}

	tokenData, err := s.authService.CreateToken(ctx, *user)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[UserService] error on CreateToken")

		return nil, err
	}

	resp := &entity.LoginResponse{
		UserID:    user.ID,
		RoleID:    user.RoleID,
		Token:     tokenData.Token,
		TokenType: tokenData.TokenType,
		ExpiredAt: tokenData.ExpiredAt,
	}

	return resp, nil
}

func (s *UserService) Register(ctx context.Context, userData entity.User) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UserService.Register")
	defer span.Finish()

	user, err := s.userRepository.GetUserByEmail(ctx, userData.Email)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on GetUserByEmail")
		return 0, err
	}

	if user.Username != "" {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on Email Already Exist")
		return 0, errors.New("Email already exist")
	}

	passwordHashed, err := utils.HashPassword(userData.Password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on HashPassword")
		return 0, errors.New("Failed hashed password")
	}

	userDataProceed := entity.UserORM{
		Username: userData.Username,
		Password: passwordHashed,
		Name:     userData.Name,
		Email:    userData.Email,
	}

	result, err := s.userRepository.RegisterUser(ctx, userDataProceed)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[UserService] error on RegisterUser")
		return 0, err
	}

	return result, nil

}
