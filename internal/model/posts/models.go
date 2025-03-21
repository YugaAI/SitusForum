package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHasta   []string `json:"postHasta"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id"`
		UserID      int64     `db:"user_id"`
		PostTitle   string    `db:"post_title"`
		PostContent string    `db:"post_content"`
		PostHasta   string    `db:"post_hasta"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string    `db:"created_by"`
		UpdatedBy   string    `db:"updated_by"`
	}
)

type (
	GetAllResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID          int64    `json:"id"`
		UserID      int64    `json:"userId"`
		Username    string   `json:"username"`
		PostTitle   string   `json:"post_title"`
		PostContent string   `json:"post_content"`
		PostHasta   []string `json:"post_hasta"`
		IsLiked     bool     `json:"isLiked"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostByIDResponse struct {
		PostDetail Post      `json:"postDetail"`
		LikeCount  int       `json:"likeCount"`
		Comment    []Comment `json:"comments"`
	}

	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"userId"`
		Username       string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)
