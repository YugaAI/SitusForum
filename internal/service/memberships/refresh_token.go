package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"main.go/internal/model/memberships"
	"main.go/pkg/jwt"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingToken, err := s.membershipsRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", err
	}

	if existingToken == nil {
		return "", errors.New("refresh token has expired")
	}

	//means the token in database is not the matched whit the request token, throw error invalid refresh token
	if existingToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipsRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not found")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
