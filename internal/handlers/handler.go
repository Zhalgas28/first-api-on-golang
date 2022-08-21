package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	// service
}

func NewHandler() *Handler {
	return &Handler{}
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
