package repository

import (
	"context"
	"go-back/internal/domain"
	postgres "go-back/internal/storage/database"
)

type UserRepository struct{}

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

func (UserRepository) getUserByUUIDQuery() string {
	return `
		SELECT uuid, name, email, created_at, updated_at
		FROM users
		WHERE uuid = $1
		LIMIT 1;
	`
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

func (u UserRepository) UpdateUser(user domain.User) (domain.User, error) {
	ctx := context.Background()
	db := postgres.GetDB()
	defer db.Close()

	_, err := db.Exec(ctx, u.updateUserQuery(), user.Name, user.Email, user.UUID)
	if err != nil {
		return domain.User{}, err
	}

	updatedUser, err := u.ListUserByUUID(user.UUID)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (UserRepository) updateUserQuery() string {
	return `
		UPDATE users
		SET name = $1,
			email = $2,
			updated_at = NOW()
		WHERE uuid = $3;
	`
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

func (UserRepository) createUserQuery() string {
	return `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING uuid, name, email, created_at, updated_at;
	`
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

func (UserRepository) deleteUserQuery() string {
	return `
		DELETE FROM users
		WHERE uuid = $1;
	`
}
