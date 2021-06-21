package user

import (
	"assesment/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllUser() ([]entity.User, error)
	FindByID(ID string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(ID string, dataUpdate map[string]interface{}) (entity.User, error)
	Delete(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(ID string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

	var user entity.User
	if err := r.db.Model(&user).Where("id=?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) Delete(ID string) (string, error) {
	var user entity.User

	if err := r.db.Where("id = ?", ID).Delete(&user).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
