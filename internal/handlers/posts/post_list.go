package posts

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.New("invalid page index").Error()})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.New("invalid page Size").Error()})
		return
	}

	response, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
