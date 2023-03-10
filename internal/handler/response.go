package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var words = map[string]int{
	"not found":     http.StatusNotFound,
	"has no values": http.StatusBadRequest,
	"another user":  http.StatusForbidden,
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	if statusCode == -1 {
		for word, code := range words {
			if strings.Contains(message, word) {
				statusCode = code
				break
			}
		}
		if statusCode == -1 {
			statusCode = http.StatusInternalServerError
		}
	}
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
