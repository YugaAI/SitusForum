package posts

import (
	"context"

	"main.go/internal/model/posts"
)

func (r *repository) GetUsersActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	query := `SELECT * FROM user_activity WHERE post_id=? AND user_id=?`

	var response posts.UserActivityModel
	row := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID)
	err := row.Scan(&response.ID, &response.PostID, &response.UserID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activity (post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activity SET is_liked=?, updated_at=?, updated_by=? WHERE post_id=? AND user_id=?`
	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.PostID, model.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) LikeCountByPostID(ctx context.Context, postID int64) (int, error) {
	query := `SELECT COUNT(id) FROM user_activity WHERE post_id=? AND is_liked= true`

	var response int
	row := r.db.QueryRowContext(ctx, query, postID)
	err := row.Scan(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}
