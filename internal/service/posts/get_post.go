package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"main.go/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostByIDResponse, error) {
	postDetail, err := s.postsRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get post by ID to database")
		return nil, err
	}

	likeCount, err := s.postsRepo.LikeCountByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get like count by post ID to database")
		return nil, err
	}

	comments, err := s.postsRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get comment by post ID to database")
		return nil, err
	}

	return &posts.GetPostByIDResponse{
		PostDetail: posts.Post{
			ID:          postDetail.ID,
			UserID:      postDetail.UserID,
			Username:    postDetail.Username,
			PostTitle:   postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHasta:   postDetail.PostHasta,
			IsLiked:     postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comment:   comments,
	}, nil
}
