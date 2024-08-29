package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSchema struct {
	gorm.Model
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
	Todos     []TodoSchema `gorm:"foreignKey:UserId"`
}
type TodoSchema struct {
	gorm.Model
	Id     uint // Standard field for the primary key
	Title  string
	Status string
	UserId uint
}
