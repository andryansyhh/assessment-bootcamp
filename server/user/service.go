package user

import (
	"assesment/entity"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() ([]UserFormatter, error)
	GetUserByID(userID string) (UserFormatter, error)
	UserRegister(input entity.UserRegister) (UserFormatter, error)
	UserLogin(input entity.UserLogin) (entity.User, error)
	UpdateUser(userID string, input entity.UpdateUser) (UserFormatter, error)
	DeleteUser(userID string) (string, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]UserFormatter, error) {
	users, err := s.repository.GetAllUser()

	var formatUsers []UserFormatter

	for _, user := range users {
		formatUser := UserFormat(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err

	}

	return formatUsers, nil
}

func (s *service) GetUserByID(userID string) (UserFormatter, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(user)

	return formatter, nil
}

func (s *service) UserRegister(input entity.UserRegister) (UserFormatter, error) {

	genPass, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser = entity.User{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(genPass),
		Address:   input.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.repository.Create(newUser)
	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(user)

	return formatter, nil
}

func (s *service) UserLogin(input entity.UserLogin) (UserLoginFormatter, error) {
	checkUser, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return UserLoginFormatter{}, err
	}

	if checkUser.ID == 0 {
		return UserLoginFormatter{}, errors.New("invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(input.Password))

	if err != nil {
		return UserLoginFormatter{}, errors.New("user email / password invalid")
	}

	formatter := UserLoginFormat(checkUser)

	return formatter, nil

}

func (s *service) UpdateUser(userID string, input entity.UpdateUser) (UserFormatter, error) {
	var dataUpdate = map[string]interface{}{}

	if input.Fullname != "" || len(input.Fullname) != 0 {
		dataUpdate["fullname"] = input.Fullname
	}

	if input.Address != "" || len(input.Address) != 0 {
		dataUpdate["address"] = input.Address
	}

	dataUpdate["updated_at"] = time.Now()

	userUpdate, err := s.repository.Update(userID, dataUpdate)

	if err != nil {
		return UserFormatter{}, err
	}

	formatUser := UserFormat(userUpdate)

	return formatUser, nil
}

func (s *service) DeleteUser(userID string) (string, error) {
	msg, err := s.repository.Delete(userID)

	if err != nil {
		return msg, err
	}

	message := fmt.Sprintf("user id %s success deleted", userID)

	return message, nil
}
