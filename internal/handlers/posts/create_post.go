package posts

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/model/posts"
)

func (h *Handler) createPost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")

	err := h.postSvc.CreatePost(ctx, userID, request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
