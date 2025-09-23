package service

import (
	"errors"
	"reflect"
	"testing"

	"go-back/internal/domain"
)

type MockUserRepository struct {
	ListAllUsersFunc   func() ([]domain.User, error)
	ListUserByUUIDFunc func(string) (domain.User, error)
}

func (m *MockUserRepository) ListAllUsers() ([]domain.User, error) {
	return m.ListAllUsersFunc()
}

func (m *MockUserRepository) ListUserByUUID(userUUID string) (domain.User, error) {
	if m.ListUserByUUIDFunc != nil {
		return m.ListUserByUUIDFunc(userUUID)
	}
	return domain.User{}, nil
}
func (m *MockUserRepository) ListUserByEmail(string) (domain.User, error) { return domain.User{}, nil }
func (m *MockUserRepository) UpdateUser(domain.User) (domain.User, error) { return domain.User{}, nil }
func (m *MockUserRepository) ManageActivateUser(string) (domain.User, error) {
	return domain.User{}, nil
}
func (m *MockUserRepository) CreateUser(domain.UserInput) (domain.User, error) {
	return domain.User{}, nil
}
func (m *MockUserRepository) DeleteUser(string) error { return nil }

func TestNewUserService(t *testing.T) {
	repo := &MockUserRepository{}
	service := NewUserService(repo)
	if service.userRepository != repo {
		t.Errorf("expected repository to be set correctly")
	}
}

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

func TestUserService_ListUserByUUID(t *testing.T) {
	t.Run("returns user when repository succeeds", func(t *testing.T) {
		mockUser := domain.User{UUID: "1", Name: "Maria", Email: "maria@example.com"}
		repo := &MockUserRepository{
			ListUserByUUIDFunc: func(userUUID string) (domain.User, error) {
				if userUUID == "1" {
					return mockUser, nil
				}
				return domain.User{}, errors.New("user not found")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.ListUserByUUID("1")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(user, mockUser) {
			t.Errorf("expected %v, got %v", mockUser, user)
		}
	})
	t.Run("returns error when repository fails", func(t *testing.T) {
		repo := &MockUserRepository{
			ListUserByUUIDFunc: func(userUUID string) (domain.User, error) {
				return domain.User{}, errors.New("db error")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.ListUserByUUID("1")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if (user != domain.User{}) {
			t.Errorf("expected empty user, got %v", user)
		}
	})
}

