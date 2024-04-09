package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) BannerGet(c *gin.Context) {
	slog.Info("handler: UserBannerGet start")
	defer slog.Info("handler: UserBannerGet end")

	tagIDs, err := getIntArrayParam(c, "tags_id")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid tag_ids param")
		return
	}

	featureId, err := getIntParam(c, "feature_id")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid feature_id param")
		return
	}

	limit, err := getDefaultIntParam(c, "limit")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid limit param")
		return
	}

	offset, err := getDefaultIntParam(c, "offset")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	banners, err := h.service.BannerGet(featureId, tagIDs, limit, offset)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Data": banners,
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
	if !exists || role != types.Admin {
		response.NewErrorResponse(c, http.StatusUnauthorized, "user role not found in context")
		return
	}

	var input types.BannerPostRequest
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.BannerPost(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) UserBannerGet(c *gin.Context) {
	slog.Info("handler: UserBannerGet start")
	defer slog.Info("handler: UserBannerGet end")
	var input types.GetModelBannerInput
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	banner, err := h.service.UserBannerGet(input.TagIds, input.FeatureId, input.UseLastRevision)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, banner)
}
