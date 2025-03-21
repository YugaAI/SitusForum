package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"main.go/internal/middlewere"
	"main.go/internal/model/memberships"
)

type membershipsService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipsService
}

func NewHandler(api *gin.Engine, membershipSvc membershipsService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middlewere.AuthRefreshMiddlewere())
	routeRefresh.POST("/refresh", h.Refresh)
}
