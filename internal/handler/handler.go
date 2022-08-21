package handler

import (
	"github.com/gin-gonic/gin"
	"man_utd/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	routes := gin.New()

	auth := routes.Group("auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}
	return routes
}
