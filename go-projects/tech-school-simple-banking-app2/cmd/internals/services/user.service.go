package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
)

type IUserService interface {
	GetAllUsers(int, int) ([]repositories.User, error)
	GetUserById(int64) (*repositories.User, error)
	CreateUser(dto.CreateAccountDto) (*dto.UserResp, error)
}

type DefaultUserService struct {
	repo repositories.IUserRepo
}

func (s *DefaultUserService) GetAllUsers(page int, pageSize int) ([]repositories.User, error) {
	return s.repo.GetAllUsers(page, pageSize)
}

func (s *DefaultUserService) GetUserById(userId int64) (*repositories.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *DefaultUserService) CreateUser(userPayload dto.CreateUserDto) (dto.UserResp, error) {
	user, stdErr := s.repo.CreateUser(userPayload)
	if stdErr != nil {
		return dto.UserResp{}, stdErr
	}

	return user.CreateUserResponse(), nil
}

func NewDefaultUserService(repo repositories.IUserRepo) *DefaultUserService {
	return &DefaultUserService{
		repo: repo,
	}
}
