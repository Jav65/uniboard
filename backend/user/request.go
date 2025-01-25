package user

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
