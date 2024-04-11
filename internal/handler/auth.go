package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param input body types.UserType true "User information"
// @Success 200 {object} map[string]interface{} "id": int
// @Failure 400 {object} response.errorResponse "Invalid input body"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
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

// SignIn godoc
// @Summary Authenticate user
// @Description Authenticate user with provided credentials
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body types.SignInInput true "Sign-in information"
// @Success 200 {object} map[string]interface{} "token": "Successful response with access token"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 401 {object} response.errorResponse "Unauthorized"
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input types.SignInInput

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.CheckAuthData(input.Username, input.Password)
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
