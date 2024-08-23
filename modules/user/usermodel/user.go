package usermodel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"not null;"`
	Email     string         `json:"email" gorm:"not null;unique;"`
	Password  *string        `json:"password" gorm:"not null;"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
