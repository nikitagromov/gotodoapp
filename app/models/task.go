package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Name string
	ProjectID uint
	OwnerID string `gorm:"ForeignKey:UserRefer"`
}
