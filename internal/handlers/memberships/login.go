package memberships

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := memberships.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(200, response)
}
