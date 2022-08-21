package handler

import (
	"github.com/gin-gonic/gin"
	"man_utd/entity"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
