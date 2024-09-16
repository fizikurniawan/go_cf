package response

import (
	"crowdfunding/internal/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success single data
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Success with pagination
func SuccessWithPagination(c *gin.Context, message string, data interface{}, page, perPage, total int) {
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Pagination: &models.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
	})
}

// Error response
func Error(c *gin.Context, message string, errors map[string][]string) {
	c.JSON(http.StatusBadRequest, models.ErrorResponse{
		Status:  "error",
		Message: message,
		Errors:  errors,
	})
}

// Internal server error
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, models.ErrorResponse{
		Status:  "error",
		Message: message,
	})
}
