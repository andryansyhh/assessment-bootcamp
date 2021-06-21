package user

import (
	"assesment/entity"
)

type UserFormatter struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type UserLoginFormatter struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Fullname      string `json:"fullname"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Authorization string `json:"authorization"`
}

func UserFormat(user entity.User) UserFormatter {
	return UserFormatter{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
	}
}

func UserLoginFomat(user entity.User) UserLoginFormatter {
	return UserLoginFormatter{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
		// Authorization: user.Address,
	}
}
