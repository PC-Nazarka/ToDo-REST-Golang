package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/internal/entity"
)

// Create Comment
//
//	@Summary		Create Comment
//	@Tags			comments
//	@Description	Create Comment
//	@ID				create-comment
//
//	@Security		ApiKeyAuth
//
// @Accept json
//
//	@Produce		json
//	@Param			input	body		entity.CommentCreate	true	"comment body"
//	@Success		200		{object}	entity.Comment
//	@Failure		400		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/comments [post]
func (h *Handler) createComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input entity.CommentCreate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	commentId, err := h.services.Comment.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	comment, err := h.services.Comment.GetById(commentId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// Update Comment Info
//
//	@Summary		Update Comment Info
//	@Tags			comments
//	@Description	Update Comment Info
//	@ID				update-comment-info
//
//	@Security		ApiKeyAuth
//
// @Accept json
//
//	@Produce		json
//	@Param			input	path		integer					true	"comment id"
//	@Param			input	body		entity.CommentUpdate	true	"updated comment info"
//	@Success		200		{object}	entity.Comment
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/comments/:id [patch]
func (h *Handler) updateComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input entity.CommentUpdate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.Comment.Update(userId, commentId, input); err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	comment, err := h.services.Comment.GetById(commentId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
}

// Delete Comment
//
//	@Summary		Delete Comment
//	@Tags			comments
//	@Description	Delete Comment
//	@ID				delete-comment
//
//	@Security		ApiKeyAuth
//
// @Accept json
//
//	@Produce		json
//	@Param			input	path	integer	true	"comment id"
//	@Success		204
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/comments/:id [delete]
func (h *Handler) deleteComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Comment.Delete(userId, commentId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
