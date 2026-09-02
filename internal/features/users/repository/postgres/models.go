package users_postgres_repository

import core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []core_domain.User {
	userDomains := make([]core_domain.User, len(users))

	for i, user := range users {
		userDomains[i] = core_domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return userDomains
}
