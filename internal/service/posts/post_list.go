package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"main.go/internal/model/posts"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.postsRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all post from database")
		return response, err
	}
	return response, nil
}
