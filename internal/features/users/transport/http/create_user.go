package users_transport_http

import (
	"net/http"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
	core_logger "github.com/Lxame/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Lxame/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Lxame/golang-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct { //validate:"required"
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"               example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+" example:"+79998887766"`
}

type CreateUserResponse UserDTOResponse

// CreateUser   godoc
// @Summary     Создать пользователя
// @Description Создать нового пользователя в системе
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request body CreateUserRequest true "CreateUser тело запроса"
// @Success     201 {object} CreateUserResponse "Успешно созданный пользователь"
// @Failure     400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /users [post]
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
