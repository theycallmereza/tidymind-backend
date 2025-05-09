package database

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Tasks    []Task    `gorm:"foreignKey:UserID"`
}

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `gorm:"not null"`
	Description *string
	Status      string `gorm:"not null;default:'pending'"`
	UserID      uuid.UUID
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
