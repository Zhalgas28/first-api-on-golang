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

	api := routes.Group("api", h.UserIdentity)
	{

		news := api.Group("/")
		{
			news.POST("/", h.CreateNews)
			news.GET("/", h.GetAllNews)
			news.GET(":id", h.NewsById)
			myNews := news.Group("my-news")
			{
				myNews.POST("/", h.CreateNews)
				myNews.GET("/", h.GetUserNews)
				myNews.GET("/:id", h.UserNewsById)
				myNews.PUT("/:id", h.UpdateNews)
				myNews.DELETE("/:id", h.DeleteNews)
			}
		}
	}

	return routes
}
