package handler

import (
	"avito_test_assingment/internal/handler/response"
	"avito_test_assingment/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BannerGet godoc
// @Summary Get banners
// @Description Get banners by tag IDs, feature ID, limit, and offset
// @Tags banners
// @Accept json
// @Produce json
// @Param tags_id query []int true "Tag IDs"
// @Param feature_id query int true "Feature ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} map[string]interface{} "Data": []Banner "Successful response with banners"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /api/v1/banners [get]
func (h *Handler) BannerGet(c *gin.Context) {
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

// BannerIdDelete godoc
// @Summary Delete a banner by ID
// @Description Delete a banner by its ID
// @Tags banners
// @Accept json
// @Produce json
// @Param id path int true "Banner ID"
// @Success 501 {object} map[string]interface{} "status": "ok" "Successful response"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /api/v1/banners/{id} [delete]
func (h *Handler) BannerIdDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.service.BannerIdDelete(id); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "ok",
	})

}

// BannerIdPatch godoc
// @Summary Update a banner by ID
// @Description Update a banner by its ID with the provided information
// @Tags banners
// @Accept json
// @Produce json
// @Param id path int true "Banner ID"
// @Param input body types.BannerIdPatchRequest true "Banner information"
// @Success 501 {object} map[string]interface{} "status": "ok" "Successful response"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /api/v1/banners/{id} [patch]
func (h *Handler) BannerIdPatch(c *gin.Context) {
	var banner types.BannerIdPatchRequest
	if err := c.BindJSON(&banner); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.service.BannerIdPatch(id, banner); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"status": "ok",
	})
}

// BannerPost godoc
// @Summary Create a new banner
// @Description Create a new banner with the provided information
// @Tags banners
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body types.BannerPostRequest true "Banner information"
// @Success 200 {object} map[string]interface{} "id": int "Successful response with banner ID"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 401 {object} response.errorResponse "Unauthorized"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /api/v1/banners [post]
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

// UserBannerGet godoc
// @Summary Get user-specific banner
// @Description Get user-specific banner by tag IDs, feature ID, and revision
// @Tags banners
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param tag_ids query []int true "Tag IDs"
// @Param feature_id query int true "Feature ID"
// @Param use_last_revision query bool false "Use last revision"
// @Success 200 {object} types.BannerGet200ResponseInner "Successful response with banner"
// @Failure 400 {object} response.errorResponse "Bad request"
// @Failure 401 {object} response.errorResponse "Unauthorized"
// @Failure 403 {object} response.errorResponse "Forbidden"
// @Failure 500 {object} response.errorResponse "Internal server error"
// @Router /api/v1/banners [get]
func (h *Handler) UserBannerGet(c *gin.Context) {
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

	role, exists := c.Get("role")
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "can't get role")
		return
	}

	if !banner.IsActive && role != types.Admin {
		response.NewErrorResponse(c, http.StatusForbidden, "banner is not active")
		return
	}

	c.JSON(http.StatusOK, banner)
}
