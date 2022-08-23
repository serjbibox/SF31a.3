package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/serjbibox/GoNews/pkg/storage"
)

type Handler struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	posts := router.Group("/posts")
	{
		posts.POST("/", h.createPost)
		posts.GET("/", h.getPosts)
		posts.PUT("/", h.updatePost)
		posts.DELETE("/:id", h.deletePost)
	}
	return router
}
