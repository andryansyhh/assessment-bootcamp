package pass

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewPassRepository(db *gorm.DB) *repository {
	return &repository{db}
}
