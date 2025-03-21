package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"main.go/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHasta := strings.Join(req.PostHasta, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:      userID,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHasta:   postHasta,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(userID, 10),
		UpdatedBy:   strconv.FormatInt(userID, 10),
	}

	err := s.postsRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create post")
		return err
	}
	return nil
}
