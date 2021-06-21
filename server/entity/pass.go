package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegister struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
