package entity

import "time"

type Pass struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Website   string    `json:"website"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
