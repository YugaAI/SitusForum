package memberships

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"main.go/internal/model/memberships"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	user = &memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
		Username:  req.Username,
	}
	return s.membershipsRepo.CreateUser(ctx, *user)
}
