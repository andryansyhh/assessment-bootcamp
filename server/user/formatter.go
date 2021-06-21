package user

import (
	"assesment/entity"
)

type Formatter struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type LoginFormatter struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Fullname      string `json:"fullname"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Authorization string `json:"authorization"`
}

func UserFormat(user entity.User) Formatter {
	return Formatter{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
	}
}

func UserLoginFomat(user entity.User) LoginFormatter {
	return LoginFormatter{
		ID:            user.ID,
		Fullname:      user.Fullname,
		Email:         user.Email,
		Address:       user.Address,
		Authorization: user.Address,
	}
}
