package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"main.go/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postsRepo.GetUsersActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivity == nil {
		if !req.IsLiked {
			return errors.New("anda belum pernah like sebeumnya")
		}
		err = s.postsRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postsRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or user activity to database")
		return err
	}

	return nil
}
