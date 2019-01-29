package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	ImageProfile string     `gorm:"column:image_profile" json:"image_profile"`
}

func (User) TableName() string {
	return "rl_users"
}
