package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	queryToken          = "token"
)

func (h *Handler) userIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			NewErrorResponse(c, http.StatusUnauthorized, "empty with header")
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			NewErrorResponse(c, http.StatusUnauthorized, "invalid with header")
			return
		}
		userId, err := h.services.JWTAuthorization.ParseToken(headerParts[1])
		if err != nil {
			NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		_, err = h.services.User.GetById(userId)
		if err != nil {
			NewErrorResponse(c, -1, err.Error())
			return
		}
		c.Set(userCtx, userId)
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "you are unauthorized")
		return 0, errors.New("you are unauthorized")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
