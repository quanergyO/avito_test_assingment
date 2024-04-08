package response

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
