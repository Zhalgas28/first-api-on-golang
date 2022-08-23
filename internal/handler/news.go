package handler

import (
	"github.com/gin-gonic/gin"
	"man_utd/entity"
	"man_utd/internal/repository"
	"net/http"
	"strconv"
)

func (h *Handler) CreateNews(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var news entity.News
	if err = c.BindJSON(&news); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newsId, err := h.service.CreateNews(userId, news)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": newsId,
	})
}

type GetAllNewsResponse struct {
	Data []entity.News `json:"data"`
}

func (h *Handler) GetAllNews(c *gin.Context) {
	news, err := h.service.GetAllNews()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllNewsResponse{
		Data: news,
	})
}

func (h *Handler) NewsById(c *gin.Context) {
	id := c.Param("id")
	newsId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id type")
		return
	}
	news, err := h.service.NewsById(newsId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"news": news,
	})
}

func (h *Handler) GetUserNews(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	news, err := h.service.GetUserNews(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllNewsResponse{
		Data: news,
	})
}

func (h *Handler) UserNewsById(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id := c.Param("id")
	newsId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id type")
		return
	}

	news, err := h.service.UserNewsById(userId, newsId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"news": news,
	})
}

func (h *Handler) UpdateNews(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id := c.Param("id")
	newsId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id type")
		return
	}

	news := repository.UpdateNewsResponse{}
	err = c.BindJSON(&news)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdateNews(userId, newsId, news)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DeleteNews(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id := c.Param("id")
	newsId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id type")
		return
	}

	err = h.service.DeleteNews(userId, newsId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
