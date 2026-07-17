package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string     `json:"id" gorm:"type:char(36);primaryKey"`
	KeycloakID  string     `json:"keycloak_id" gorm:"type:varchar(36);uniqueIndex;not null"`
	Username    string     `json:"username" gorm:"type:varchar(255);not null"`
	FirstName   string     `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName    string     `json:"last_name" gorm:"type:varchar(255);not null"`
	Email       string     `json:"email" gorm:"type:varchar(255);index;not null"`
	LastLoginAt *time.Time `json:"last_login_at" gorm:"type:datetime"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
