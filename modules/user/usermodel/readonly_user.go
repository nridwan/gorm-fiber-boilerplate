package usermodel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReadonlyUserModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"not null;"`
	Email     string         `json:"email" gorm:"not null;unique;"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (ReadonlyUserModel) TableName() string {
	return "users"
}
