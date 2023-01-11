package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo-list/docs"

	"todo-list/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//	@title			ToDoListApi
//	@version		1.0
//	@description	API Server for TodoApp

//	@host		localhost:8000
//	@BasePath	/api

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/assets", "./assets")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		token := api.Group("/token")
		{
			token.POST("/create", h.createToken)
			token.POST("/refresh", h.refreshTokens)
		}
		users := api.Group("/users")
		{
			users.POST("", h.createUser)
			users.GET("/:id", h.userIdentity(), h.getUserById)
			users.GET("/me", h.userIdentity(), h.getSelfUser)
			users.PATCH("/:id", h.userIdentity(), h.updateUser)
			users.DELETE("/:id", h.userIdentity(), h.deleteUser)
			users.GET("/:id/tasks", h.userIdentity(), h.getTasksByUserId)
			users.GET("/:id/posts", h.userIdentity(), h.getPostsByUserId)
		}
		todo := api.Group("/tasks", h.userIdentity())
		{
			todo.POST("", h.createTask)
			todo.GET("/:id", h.getTaskById)
			todo.PATCH("/:id", h.updateTask)
			todo.DELETE("/:id", h.deleteTask)
			todo.POST("/import", h.importTasks)
			//todo.GET("/export", h.exportTasks)
		}
		posts := api.Group("/posts")
		{
			posts.POST("", h.userIdentity(), h.createPost)
			posts.GET("", h.getAllPosts)
			posts.GET("/:id/comments", h.getCommentsByPostId)
			posts.PATCH("/:id", h.userIdentity(), h.updatePost)
			posts.DELETE("/:id", h.userIdentity(), h.deletePost)
		}
		comment := api.Group("/comments", h.userIdentity())
		{
			comment.POST("", h.createComment)
			comment.PATCH("/:id", h.updateComment)
			comment.DELETE("/:id", h.deleteComment)
		}
	}
	return router
}
