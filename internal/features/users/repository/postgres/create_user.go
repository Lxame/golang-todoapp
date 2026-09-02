package users_postgres_repository

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user core_domain.User,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.users (full_name, phone_number)
		VALUES ($1, $2)
		RETURNING id, version, full_name, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := core_domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
