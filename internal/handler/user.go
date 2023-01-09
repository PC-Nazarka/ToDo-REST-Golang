package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/internal/entity"
)

// Create User
//
//	@Summary		Create User
//	@Tags			users
//	@Description	Create User
//	@ID				create-user
//
// Accept json
//
//	@Produce		json
//	@Param			input	body		entity.UserCreate	true	"user body"
//	@Success		200		{object}	entity.User
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input entity.UserCreate
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.User.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := h.services.User.GetById(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Get User Info
//
//	@Summary		Get User Info
//	@Tags			users
//	@Description	Get User Info
//	@ID				get-users-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer	true	"user id"
//	@Success		200		{object}	entity.User
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/users/:id [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	user, err := h.services.User.GetById(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// Get Self User Info
//
//	@Summary		Get Self User Info
//	@Tags			users
//	@Description	Get Self User Info
//	@ID				get-self-users-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Success		200	{object}	entity.User
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/me [get]
func (h *Handler) getSelfUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	user, err := h.services.User.GetById(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update User Info
//
//	@Summary		Update User Info
//	@Tags			users
//	@Description	Update User Info
//	@ID				update-user-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer				true	"user id"
//	@Param			input	body		entity.UserUpdate	true	"updated user info"
//	@Success		200		{object}	entity.User
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/users/:id [patch]
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	userPathId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if userId != userPathId {
		NewErrorResponse(c, http.StatusForbidden, "you can't update another user")
		return
	}
	var input entity.UserUpdate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.User.Update(userId, input); err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	user, err := h.services.User.GetById(userPathId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// Delete User
//
//	@Summary		Delete User
//	@Tags			users
//	@Description	Delete User
//	@ID				delete-user
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path	integer	true	"user id"
//	@Success		204
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/:id [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	userPathId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if userId != userPathId {
		NewErrorResponse(c, http.StatusForbidden, "you can't delete another user")
		return
	}
	if err = h.services.User.Delete(userId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Get Tasks of User
//
//	@Summary		Get Tasks of User
//	@Tags			users
//	@Description	Get Tasks of User
//	@ID				get-tasks-user
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer	true	"user id"
//	@Success		200		{array}		entity.Task
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/users/:id/tasks [get]
func (h *Handler) getTasksByUserId(c *gin.Context) {
	userPathId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	_, err = h.services.User.GetById(userPathId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	tasks, err := h.services.Task.GetByUserId(userPathId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if tasks == nil {
		tasks = make([]entity.Task, 0)
	}
	c.JSON(http.StatusOK, tasks)
}

// Get Posts of User
//
//	@Summary		Get Posts of User
//	@Tags			users
//	@Description	Get Posts of User
//	@ID				get-posts-user
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer	true	"user id"
//	@Success		200		{array}		entity.Post
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/users/:id/posts [get]
func (h *Handler) getPostsByUserId(c *gin.Context) {
	userPathId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	_, err = h.services.User.GetById(userPathId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	posts, err := h.services.Post.GetByUserId(userPathId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	if posts == nil {
		posts = make([]entity.Post, 0)
	}
	c.JSON(http.StatusOK, posts)
}
