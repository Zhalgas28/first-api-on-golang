package handler

import (
	"github.com/gin-gonic/gin"
	"man_utd/pkg/logging"
)

func NewErrorResponse(c *gin.Context, code int, message string) {
	logger := logging.GetLogger()
	logger.Error(message)
	c.AbortWithStatusJSON(code, gin.H{message: message})
}
