package handler

import (
	_ "avito_test_assingment/docs"
	"avito_test_assingment/internal/cache"
	"avito_test_assingment/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
	cache   cache.Cache
}

func NewHandler(service *service.Service) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.userIdentity)
	{
		banner := api.Group("/banner", h.administratorVerification)
		{
			banner.GET("/", h.BannerGet)
			banner.DELETE("/:id", h.BannerIdDelete)
			banner.PATCH("/:id", h.BannerIdPatch)
			banner.POST("/", h.BannerPost)

		}
		api.GET("/user_banner", h.UserBannerGet)
	}

	return router
}
