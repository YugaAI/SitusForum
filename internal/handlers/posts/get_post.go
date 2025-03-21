package posts

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()
	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.New("invalid post id").Error()})
		return
	}

	response, err := h.postSvc.GetPostByID(ctx, postID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}
