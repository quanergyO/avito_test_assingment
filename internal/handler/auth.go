package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user types.UserType
	if err := c.BindJSON(&user); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input types.SignInInput

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.CheckAuthData(input.Username, input.Password)
	slog.Info("UserRole for signIn: ", user.Role)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
