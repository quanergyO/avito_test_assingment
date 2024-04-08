package handler

import (
	"avito_test_assingment/internal/handler/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) BannerGet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "Not implemented",
	})
}

func (h *Handler) BannerIdDelete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "Not implemented",
	})

}

func (h *Handler) BannerIdPatch(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "Not implemented",
	})
}

func (h *Handler) BannerPost(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "Роль пользователя не найдена в контексте")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": role,
	})

}

func (h *Handler) UserBannerGet(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "Not implemented",
	})
}
