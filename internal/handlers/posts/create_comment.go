package posts

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model/posts"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreateCommentRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postID, userID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
