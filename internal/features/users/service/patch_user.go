package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch core_domain.UserPatch,
) (core_domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return core_domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, nil
}
