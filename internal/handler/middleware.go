package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	logrus.Info(userId)
	c.Set("userId", userId)
}
