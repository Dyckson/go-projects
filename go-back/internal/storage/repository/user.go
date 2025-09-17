package repository

import (
	"context"
	"go-back/internal/domain"
	postgres "go-back/internal/storage/database"
)

type UserRepository struct{}

func (u UserRepository) ListAllUsers() ([]domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var users []domain.User
	err := db.Query(ctx, &users, u.getAllUsersQuery())
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) ListUserByUUID(userUUID string) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var user domain.User
	err := db.QueryOne(ctx, &user, u.getUserByUUIDQuery(), userUUID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) ListUserByEmail(email string) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var user domain.User
	err := db.QueryOne(ctx, &user, u.getUserByEmailQuery(), email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (UserRepository) getUserByEmailQuery() string {
	return `
		SELECT uuid, name, email, created_at, updated_at
		FROM users
		WHERE email = $1
		LIMIT 1;
	`
}

func (u UserRepository) ManageActivateUser(userUUID string) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var updatedUser domain.User
	err := db.QueryOne(ctx, &updatedUser, u.manageActivateUserQuery(), userUUID)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (u UserRepository) UpdateUser(user domain.User) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var updatedUser domain.User
	err := db.QueryOne(ctx, &updatedUser, u.updateUserQuery(),
		user.Name, user.Email, user.UUID)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (u UserRepository) CreateUser(user domain.UserInput) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	var createdUser domain.User
	err := db.QueryOne(ctx, &createdUser, u.createUserQuery(), user.Name, user.Email)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

func (u UserRepository) DeleteUser(userUUID string) error {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	_, err := db.Exec(ctx, u.deleteUserQuery(), userUUID)
	if err != nil {
		return err
	}

	return nil
}

func (UserRepository) getAllUsersQuery() string {
	return `
		SELECT uuid, name, email, created_at, updated_at, is_active
		FROM users
	`
}

func (UserRepository) getUserByUUIDQuery() string {
	return `
		SELECT uuid, name, email, created_at, updated_at, is_active
		FROM users
		WHERE uuid = $1
		LIMIT 1;
	`
}

func (UserRepository) createUserQuery() string {
	return `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING uuid, name, email, created_at, updated_at, is_active;
	`
}

func (UserRepository) updateUserQuery() string {
	return `
		UPDATE users
		SET name = $1,
			email = $2,
			updated_at = NOW()
		WHERE uuid = $3
		RETURNING uuid, name, email, created_at, updated_at, is_active;
	`
}

func (UserRepository) manageActivateUserQuery() string {
	return `
		UPDATE users
		SET is_active = NOT is_active,
		    updated_at = NOW()
		WHERE uuid = $1
		RETURNING uuid, name, email, created_at, updated_at, is_active;
	`
}

func (UserRepository) deleteUserQuery() string {
	return `
		DELETE FROM users
		WHERE uuid = $1;
	`
}
