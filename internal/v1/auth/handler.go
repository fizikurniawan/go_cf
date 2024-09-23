// internal/v1/auth/handler.go
package auth

import (
	"crowdfunding/internal/common/utils"
	"crowdfunding/internal/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

// NewHandler creates a new auth handler
func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var dto RegisterUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		fieldErrors := utils.FormatValidationErrors(err, &dto) // Gunakan fungsi utilitas

		response.Error(c, "Invalid input", fieldErrors)
		return
	}

	// validate user has already exists or not by email
	user, err := h.Service.GetByEmail(dto.Email)
	if err != nil {
		response.Error(c, "Failed to register user", map[string][]string{"non_field_error": {err.Error()}})
		return
	}
	if user.ID.String() != "" {
		response.Error(c, "Failed to register user", map[string][]string{"email": {"Email has already exists"}})
		return
	}

	err = h.Service.RegisterUser(dto)
	if err != nil {
		response.InternalServerError(c, "Failed to register user")
		return
	}

	response.Success(c, "User registered successfully", nil)
}

func (h *Handler) Login(c *gin.Context) {
	var dto LoginUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		fieldErrors := utils.FormatValidationErrors(err, &dto) // Gunakan fungsi utilitas

		response.Error(c, "Invalid input", fieldErrors)
		return
	}

	user, token, refresh, err := h.Service.Login(dto.Email, dto.Password)
	if err != nil {
		response.Error(c, "Failed to Login user", map[string][]string{"non_field_error": {err.Error()}})
		return
	}

	data := map[string]string{
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"token":         token,
		"refresh_token": refresh,
	}

	response.Success(c, "Loggin successfully", data)

}
