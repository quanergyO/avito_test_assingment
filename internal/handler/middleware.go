package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
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

func (h *Handler) AdministratorVerification(c *gin.Context) {
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

func getDefaultIntParam(c *gin.Context, paramString string) (int, error) {
	param := c.DefaultQuery(paramString, "0")
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return paramInt, nil
}

func getIntParam(c *gin.Context, paramString string) (int, error) {
	param := c.Query(paramString)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return paramInt, nil
}

func getIntArrayParam(c *gin.Context, paramString string) ([]int, error) {
	items := c.QueryArray(paramString)
	itemsInt := make([]int, 0, len(items))
	for _, item := range items {
		tagID, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		itemsInt = append(itemsInt, tagID)
	}

	return itemsInt, nil
}
