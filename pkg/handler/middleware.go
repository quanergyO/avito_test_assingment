package handler

import (
	"avito_test_assingment/pkg/handler/response"
	"github.com/gin-gonic/gin"
	"log/slog"
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
	slog.Info("ROLE:", claims.Role)
	c.Set("role", claims.Role)

}
