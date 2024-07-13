package service

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
	"github.com/voltgizerz/POS-restaurant/internal/utils"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type AuthService struct {
	jwtService     ports.IJWTAuth
	txRepository   ports.ITxRepository
	userRepository ports.IUserRepository
}

func NewAuthService(i interactor.AuthService) *AuthService {
	return &AuthService{
		jwtService:     i.JWTService,
		txRepository:   i.TxRepository,
		userRepository: i.UserRepository,
	}
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (*entity.LoginResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.AuthService.Login")
	defer span.Finish()

	user, err := s.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[AuthService] error on GetUserByUsername")

		return nil, err
	}

	err = utils.VerifyPassword(password, user.PasswordHashed)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[AuthService] error on VerifyPassword")

		return nil, err
	}

	tokenData, err := s.jwtService.CreateToken(ctx, *user)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("[AuthService] error on CreateToken")

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

func (s *AuthService) Register(ctx context.Context, userData entity.User) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.AuthService.Register")
	defer span.Finish()

	user, err := s.userRepository.GetUserByEmail(ctx, userData.Email)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on GetUserByEmail")
		return 0, err
	}

	if user.Username != "" {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on Email Already Exist")
		return 0, errors.New("email already exist")
	}

	passwordHashed, err := utils.HashPassword(userData.Password)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on HashPassword")
		return 0, errors.New("failed hashed password")
	}

	userDataProceed := models.UserORM{
		Username:       userData.Username,
		PasswordHashed: passwordHashed,
		Name:           userData.Name,
		Email:          userData.Email,
	}

	tx, err := s.txRepository.StartTx(ctx)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on StartTx")
		return 0, err
	}
	defer s.txRepository.RollbackTx(ctx, tx)

	result, err := s.userRepository.RegisterUser(ctx, tx, userDataProceed)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on RegisterUser")
		return 0, err
	}

	err = s.txRepository.CommitTx(ctx, tx)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"username": userData.Username,
			"error":    err,
		}).Error("[AuthService] error on CommitTx")
		return 0, err
	}

	return result, nil
}
