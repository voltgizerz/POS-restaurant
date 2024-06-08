package service

import (
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type UserService struct {
	userRepository ports.IUserRepository
}

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	// TODO
	return nil
}

func (s *UserService) Register(email string, password string, confirmPass string) error {
	// TODO
	return nil
}
