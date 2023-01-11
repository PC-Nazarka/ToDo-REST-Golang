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
	AccessToken  string `json:"access" binding:"required"`
	RefreshToken string `json:"refresh" binding:"required"`
}

type refreshTokens struct {
	RefreshToken string `json:"token" binding:"required"`
}

// Login
//
//	@Summary		Login
//	@Tags			auth
//	@Description	Login in account
//	@ID				login-account
//
// @Accept json
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
	accessToken, refreshToken, err := h.services.JWTAuthorization.GenerateAccessRefreshTokens(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, signInOutput{accessToken, refreshToken})
}

// Refresh
//
//	@Summary		Refresh
//	@Tags			auth
//	@Description	Refresh tokens
//	@ID				refresh-tokens
//
// @Accept json
//
//	@Produce		json
//	@Param			input	body		refreshTokens	true	"refresh token"
//	@Success		200		{object}	signInOutput
//	@Failure		400		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/refresh [post]
func (h *Handler) refreshTokens(c *gin.Context) {
	var input refreshTokens
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.JWTAuthorization.ParseToken(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, err := h.services.JWTAuthorization.GenerateAccessRefreshTokens(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, signInOutput{accessToken, refreshToken})
}
