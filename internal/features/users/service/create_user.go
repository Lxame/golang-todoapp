package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user core_domain.User,
) (core_domain.User, error) {
	if err := user.Validate(); err != nil {
		return core_domain.User{}, nil
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
