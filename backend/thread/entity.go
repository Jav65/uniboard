package thread

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	ID        int            `json:"id"`
	AuthorID  int            `json:"author_id"`
	Title     string         `json:"title" binding:"required"`
	Content   string         `json:"content" binding:"required"`
	Likes     int            `json:"likes"`
	Comments  int            `json:"comments"`
	Views     int            `json:"views"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
