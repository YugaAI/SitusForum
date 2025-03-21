package posts

import (
	"context"

	"main.go/internal/configs"
	"main.go/internal/model/posts"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUsersActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllResponse, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)

	LikeCountByPostID(ctx context.Context, postID int64) (int, error)

	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)
}

type service struct {
	cfg       *configs.Config
	postsRepo postsRepository
}

func NewService(cfg *configs.Config, postsRepo postsRepository) *service {
	return &service{
		cfg:       cfg,
		postsRepo: postsRepo,
	}
}
