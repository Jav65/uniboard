package thread

import "time"

type ThreadRequest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Author_id   int       `json:"author_id"`
	Author_name string    `json:"author_name"`
	Tags_name   []string  `json:"tags_name"`
	Tags        []int     `json:"tags"`
	Likes       int       `json:"likes"`
	Comments    int       `json:"comments"`
	Views       int       `json:"views"`
	CreatedAt  time.Time `json:"created_at"`
}
