package service

import (
	"errors"
	"reflect"
	"testing"

	"go-back/internal/domain"
)

var mockUser = domain.User{UUID: "1", Name: "John", Email: "john@example.com"}

type MockUserRepository struct {
	ListAllUsersFunc    func() ([]domain.User, error)
	ListUserByUUIDFunc  func(string) (domain.User, error)
	ListUserByEmailFunc func(string) (domain.User, error)
	UpdateUserFunc      func(domain.User) (domain.User, error)
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

func (m *MockUserRepository) ListUserByEmail(email string) (domain.User, error) {
	if m.ListUserByEmailFunc != nil {
		return m.ListUserByEmailFunc(email)
	}
	return domain.User{}, nil
}

func (m *MockUserRepository) UpdateUser(u domain.User) (domain.User, error) {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(u)
	}
	return domain.User{}, nil
}
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
		repo := &MockUserRepository{
			ListAllUsersFunc: func() ([]domain.User, error) {
				return []domain.User{mockUser}, nil
			},
		}
		service := UserService{userRepository: repo}
		users, err := service.ListAllUsers()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(users, []domain.User{mockUser}) {
			t.Errorf("expected %v, got %v", []domain.User{mockUser}, users)
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
		repo := &MockUserRepository{
			ListUserByUUIDFunc: func(userUUID string) (domain.User, error) {
				if userUUID == mockUser.UUID {
					return mockUser, nil
				}
				return domain.User{}, errors.New("user not found")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.ListUserByUUID(mockUser.UUID)
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

func TestUserService_ListUserByEmail(t *testing.T) {
	t.Run("returns user when repository succeeds", func(t *testing.T) {
		repo := &MockUserRepository{
			ListUserByEmailFunc: func(email string) (domain.User, error) {
				if email == mockUser.Email {
					return mockUser, nil
				}
				return domain.User{}, errors.New("user not found")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.ListUserByEmail(mockUser.Email)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(user, mockUser) {
			t.Errorf("expected %v, got %v", mockUser, user)
		}
	})
	t.Run("returns error when repository fails", func(t *testing.T) {
		repo := &MockUserRepository{
			ListUserByEmailFunc: func(email string) (domain.User, error) {
				return domain.User{}, errors.New("db error")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.ListUserByEmail(mockUser.Email)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if (user != domain.User{}) {
			t.Errorf("expected empty user, got %v", user)
		}
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	t.Run("returns updated user when repository succeeds", func(t *testing.T) {
		repo := &MockUserRepository{
			UpdateUserFunc: func(u domain.User) (domain.User, error) {
				return mockUser, nil
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.UpdateUser(mockUser)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(user, mockUser) {
			t.Errorf("expected %v, got %v", mockUser, user)
		}
	})
	t.Run("returns error when repository fails", func(t *testing.T) {
		repo := &MockUserRepository{
			UpdateUserFunc: func(u domain.User) (domain.User, error) {
				return domain.User{}, errors.New("db error")
			},
		}
		service := UserService{userRepository: repo}
		user, err := service.UpdateUser(mockUser)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if user != (domain.User{}) {
			t.Errorf("expected empty user, got %v", user)
		}
	})
}
