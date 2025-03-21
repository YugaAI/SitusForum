package posts

import (
	"context"
	"strings"

	"main.go/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO post (user_id, post_title, post_content, post_hasta, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHasta, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllResponse, error) {
	query := ` SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hasta FROM post p JOIN users u ON p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	var response posts.GetAllResponse

	row, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer row.Close()

	data := make([]posts.Post, 0)
	for row.Next() {
		var (
			models   posts.PostModel
			username string
		)
		err = row.Scan(&models.ID, &models.UserID, &username, &models.PostTitle, &models.PostContent, &models.PostHasta)
		if err != nil {
			return response, err
		}
		data = append(data, posts.Post{
			ID:          models.ID,
			UserID:      models.UserID,
			Username:    username,
			PostTitle:   models.PostTitle,
			PostContent: models.PostContent,
			PostHasta:   strings.Split(models.PostHasta, ","),
		})
	}

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return response, nil
}

func (r *repository) GetPostByID(ctx context.Context, id int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hasta, uv.is_liked FROM post p JOIN users u ON p.user_id = u.id JOIN user_activity uv ON uv.post_id = p.id WHERE p.id = ? `

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHasta, &isLiked)
	if err != nil {
		return nil, err
	}
	return &posts.Post{
		ID:          model.ID,
		UserID:      model.UserID,
		Username:    username,
		PostTitle:   model.PostTitle,
		PostContent: model.PostContent,
		PostHasta:   strings.Split(model.PostHasta, ","),
		IsLiked:     isLiked,
	}, nil

}
