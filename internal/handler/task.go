package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/internal/entity"
)

// Create Task
//
//	@Summary		Create Task
//	@Tags			tasks
//	@Description	Create task
//	@ID				create-task
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	body		entity.TaskCreate	true	"task body"
//	@Success		200		{object}	entity.Task
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input entity.TaskCreate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	taskId, err := h.services.Task.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	task, err := h.services.Task.GetById(taskId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Get Task Info
//
//	@Summary		Get Task Info
//	@Tags			tasks
//	@Description	Get Task Info
//	@ID				get-task-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer	true	"task id"
//	@Success		200		{object}	entity.Task
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/tasks/:id [get]
func (h *Handler) getTaskById(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	task, err := h.services.Task.GetById(taskId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

// Update Task Info
//
//	@Summary		Update Task Info
//	@Tags			tasks
//	@Description	Update Task Info
//	@ID				update-task-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer				true	"task id"
//	@Param			input	body		entity.TaskUpdate	true	"updated task info"
//	@Success		200		{object}	entity.Task
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/tasks/:id [patch]
func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input entity.TaskUpdate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.Task.Update(userId, taskId, input); err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	task, err := h.services.Task.GetById(taskId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

// Delete Task
//
//	@Summary		Delete Task
//	@Tags			tasks
//	@Description	Delete Task
//	@ID				delete-task
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path	integer	true	"task id"
//	@Success		204
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/tasks/:id [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err = h.services.Task.Delete(userId, taskId); err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
