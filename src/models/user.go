package models

import (
	"time"
)

type User struct {
	ID           uint       `gorm:"column:id;primary_key" json:"id"`
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	ImageProfile string     `gorm:"column:image_profile" json:"image_profile"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "rl_users"
}
