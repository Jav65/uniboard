package comment

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Thread_id int       `json:"thread_id"`
	Content   string    `json:"content"`
	Author_id    int    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
