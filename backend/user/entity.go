package user

import "time"

type User struct {
	Id        int       `json:"id"`
	Username      string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
