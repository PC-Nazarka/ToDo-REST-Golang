package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
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
// @Accept json
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
// @Accept json
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
// @Accept json
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
// @Accept json
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

// Import Task
//
//	@Summary		Import Task
//	@Tags			tasks
//	@Description	Import Task
//	@ID				import-task
//
//	@Security		ApiKeyAuth
//
// @Accept multipart/form-data
//
//	@Produce		json
//	@Param			file formData file true "CSV file with tasks"
//	@Success		201
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/tasks/import [post]
func (h *Handler) importTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	filename := strings.Join(strings.Split(file.Filename, " "), "_")
	path := "assets/upload/" + filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !strings.Contains(filename, ".csv") {
		NewErrorResponse(c, http.StatusBadRequest, "assets file has invalid type")
		return
	}
	tasks, err := h.services.Task.ParseFile(path)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	ids, err := h.services.Task.CreateBulk(userId, tasks)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	createdTasks, err := h.services.Task.GetByIds(ids)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	os.Remove(path)
	c.JSON(http.StatusCreated, createdTasks)
}

// Export Task
//
//	@Summary		Export Task
//	@Tags			tasks
//	@Description	Export Task
//	@ID				Export-task
//
//	@Security		ApiKeyAuth
//
//	@Produce		csv
//	@Success		200
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/tasks/export [get]
func (h *Handler) exportTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	tasks, err := h.services.Task.GetByUserId(userId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	filename := fmt.Sprintf("tasks_%d.csv", userId)
	path := "assets/download/" + filename
	err = h.services.Task.WriteFile(tasks, path)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/CSV")
	c.File(path)
	os.Remove(path)
}
