package comment

import "time"

type CommentRequest struct {
	ID        int       `json:"id"`
	Thread_id int       `json:"thread_id"`
	Content   string    `json:"content"`
	Author_id int       `json:"author_id"`
	Author_name string `json:"author_name"`
	CreatedAt time.Time `json:"created_at"`
}
