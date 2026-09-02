package users_transport_http

import (
	"net/http"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
	core_logger "github.com/Lxame/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Lxame/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Lxame/golang-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct { //validate:"required"
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responserHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateUser handler")

	var requset CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &requset); err != nil {
		responserHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain, err := h.usersService.CreateUser(ctx, domainFromDTO(requset))
	if err != nil {
		responserHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))
	responserHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) core_domain.User {
	return core_domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
