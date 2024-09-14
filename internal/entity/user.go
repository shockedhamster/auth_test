package entity

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
