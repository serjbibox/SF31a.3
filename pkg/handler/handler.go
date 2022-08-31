package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/GoNews/pkg/storage"
)

// Обработчик HTTP запросов сервера GoNews
type Handler struct {
	storage *storage.Storage
}

//Конструктор объекта Handler
func New(storage *storage.Storage) (*Handler, error) {
	if storage == nil {
		return nil, errors.New("storage is nil")
	}
	return &Handler{storage: storage}, nil
}

//Инициализация маршрутизатора запросов.
//Регистрация обработчиков запросов
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
