package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"main.go/internal/middlewere"
	"main.go/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostByIDResponse, error)
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middlewere.AuthMiddlewere())

	route.POST("/create", h.createPost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:postID", h.GetPostByID)
}
