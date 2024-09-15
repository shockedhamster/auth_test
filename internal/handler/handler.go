package handler

import (
	_ "github.com/auth_test/cmd/docs"
	"github.com/auth_test/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	testAPI := router.Group("/api", h.userIdentity)
	{
		testAPI.POST("/hello", h.hello)
		testAPI.DELETE("/delete-user", h.deleteUser) // вынести код по удалению в отдельный сервис
	}

	return router
}
