package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/internal/entity"
)

// Create Post
//
//	@Summary		Create Post
//	@Tags			posts
//	@Description	Create Post
//	@ID				create-posts
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	body		entity.PostCreate	true	"post body"
//	@Success		200		{object}	entity.Post
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/posts [post]
func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input entity.PostCreate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	postId, err := h.services.Post.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	post, err := h.services.Post.GetById(postId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusCreated, post)
}

// Get All Posts
//
//	@Summary		Get All Posts
//	@Tags			posts
//	@Description	Get All Posts
//	@ID				get-all-posts
//
// Accept json
//
//	@Produce		json
//	@Success		200	{array}		entity.Post
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/posts [get]
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Post.GetAll()
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	if posts == nil {
		posts = make([]entity.Post, 0)
	}
	c.JSON(http.StatusOK, posts)
}

// Update Post Info
//
//	@Summary		Update Post Info
//	@Tags			posts
//	@Description	Update Post Info
//	@ID				update-post-info
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer				true	"post id"
//	@Param			input	body		entity.PostUpdate	true	"updated post info"
//	@Success		200		{object}	entity.Post
//	@Failure		400		{object}	errorResponse
//	@Failure		401		{object}	errorResponse
//	@Failure		403		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/posts/:id [patch]
func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input entity.PostUpdate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Post.Update(userId, postId, input)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	post, err := h.services.Post.GetById(postId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

// Delete Post
//
//	@Summary		Delete Post
//	@Tags			posts
//	@Description	Delete Post
//	@ID				delete-post
//
//	@Security		ApiKeyAuth
//
// Accept json
//
//	@Produce		json
//	@Param			input	path	integer	true	"post id"
//	@Success		204
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		403	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/posts/:id [delete]
func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Post.Delete(userId, postId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Get Comments of Post
//
//	@Summary		Get Comments of Post
//	@Tags			posts
//	@Description	Get Comments of Post
//	@ID				get-comments-post
//
// Accept json
//
//	@Produce		json
//	@Param			input	path		integer	true	"post id"
//	@Success		200		{array}		entity.Comment
//	@Failure		400		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/api/posts/:id/comments [get]
func (h *Handler) getCommentsByPostId(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	comments, err := h.services.Comment.GetByPostId(postId)
	if err != nil {
		NewErrorResponse(c, -1, err.Error())
		return
	}
	if comments == nil {
		comments = make([]entity.Comment, 0)
	}
	c.JSON(http.StatusOK, comments)
}
