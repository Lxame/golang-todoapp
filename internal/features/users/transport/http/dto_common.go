package users_transport_http

import core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"           example:"10"`
	Version     int     `json:"version"      example:"1"`
	FullName    string  `json:"full_name"    example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+79998887766"`
}

func userDTOFromDomain(user core_domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []core_domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
