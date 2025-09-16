package service

import (
	"go-back/internal/domain"
	"go-back/internal/storage/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
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
