package service

import (
	"errors"
	"reflect"
	"testing"

	"go-back/internal/domain"
)

type MockUserRepository struct {
	ListAllUsersFunc func() ([]domain.User, error)
}

func (m *MockUserRepository) ListAllUsers() ([]domain.User, error) {
	return m.ListAllUsersFunc()
}

func (m *MockUserRepository) ListUserByUUID(string) (domain.User, error)  { return domain.User{}, nil }
func (m *MockUserRepository) ListUserByEmail(string) (domain.User, error) { return domain.User{}, nil }
func (m *MockUserRepository) UpdateUser(domain.User) (domain.User, error) { return domain.User{}, nil }
func (m *MockUserRepository) ManageActivateUser(string) (domain.User, error) {
	return domain.User{}, nil
}
func (m *MockUserRepository) CreateUser(domain.UserInput) (domain.User, error) {
	return domain.User{}, nil
}
func (m *MockUserRepository) DeleteUser(string) error { return nil }

func TestUserService_ListAllUsers(t *testing.T) {
	t.Run("returns users when repository succeeds", func(t *testing.T) {
		mockUsers := []domain.User{{UUID: "1", Name: "Alice", Email: "alice@example.com"}}
		repo := &MockUserRepository{
			ListAllUsersFunc: func() ([]domain.User, error) {
				return mockUsers, nil
			},
		}
		service := UserService{userRepository: repo}
		users, err := service.ListAllUsers()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(users, mockUsers) {
			t.Errorf("expected %v, got %v", mockUsers, users)
		}
	})

	t.Run("returns error when repository fails", func(t *testing.T) {
		repo := &MockUserRepository{
			ListAllUsersFunc: func() ([]domain.User, error) {
				return nil, errors.New("db error")
			},
		}
		service := UserService{userRepository: repo}
		users, err := service.ListAllUsers()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if users == nil || len(users) != 0 {
			t.Errorf("expected empty users slice, got %v", users)
		}
	})
}
