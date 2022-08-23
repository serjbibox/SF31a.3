package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/GoNews/pkg/models"
)

func (h *Handler) createPost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Post.Create(post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPosts(c *gin.Context) {
	posts, err := h.storage.Post.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) updatePost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.storage.Post.Update(post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": post.ID,
	})
}

func (h *Handler) deletePost(c *gin.Context) {
	id := c.Param("id")
	err := h.storage.Post.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
