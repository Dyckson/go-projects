package service

import (
	"go-back/internal/domain"
)

type UserRepository interface {
	ListAllUsers() ([]domain.User, error)
	ListUserByUUID(string) (domain.User, error)
	ListUserByEmail(string) (domain.User, error)
	UpdateUser(domain.User) (domain.User, error)
	ManageActivateUser(string) (domain.User, error)
	CreateUser(domain.UserInput) (domain.User, error)
	DeleteUser(string) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		userRepository: repo,
	}
}

func (us UserService) ListAllUsers() ([]domain.User, error) {
	users, err := us.userRepository.ListAllUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (us UserService) ListUserByUUID(userUUID string) (domain.User, error) {
	user, err := us.userRepository.ListUserByUUID(userUUID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) ListUserByEmail(email string) (domain.User, error) {
	user, err := us.userRepository.ListUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) UpdateUser(user domain.User) (domain.User, error) {
	user, err := us.userRepository.UpdateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) ManageActivateUser(userUUID string) (domain.User, error) {
	user, err := us.userRepository.ManageActivateUser(userUUID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) CreateUser(user domain.UserInput) (domain.User, error) {
	createdUser, err := us.userRepository.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return createdUser, nil
}

func (us UserService) DeleteUser(userUUID string) error {
	err := us.userRepository.DeleteUser(userUUID)
	if err != nil {
		return err
	}
	return nil
}
