package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signInOutput struct {
	Token string `json:"token" binding:"required"`
}

// Login
//
//	@Summary		Login
//	@Tags			auth
//	@Description	Login in account
//	@ID				login-account
//
// Accept json
//
//	@Produce		json
//	@Param			input	body		signInInput	true	"account login"
//	@Success		200		{object}	signInOutput
//	@Failure		400		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/token [post]
func (h *Handler) createToken(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.User.GetByUsernameAndPassword(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	token, err := h.services.JWTAuthorization.GenerateToken(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, signInOutput{token})
}
