package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
	core_errors "github.com/Lxame/golang-todoapp/internal/core/errors"
	core_postgres_pool "github.com/Lxame/golang-todoapp/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user core_domain.User,
) (core_domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.users
		SET
			full_name=$1,
			phone_number=$2,
			version=version+1
		WHERE id=$3 AND version=$4
		RETURNING id, version, full_name, phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.FullName,
		user.PhoneNumber,
		id,
		user.Version,
	)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return core_domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return core_domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return core_domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	), nil
}
