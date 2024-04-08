package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	claims, err := h.service.ParserToken(headerParts[1])
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("role", claims.Role)
}

func (h *Handler) administratorVerification(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "can't get role")
		return
	}

	if role != types.Admin {
		response.NewErrorResponse(c, http.StatusForbidden, "You are not an admin")
		return
	}
}
