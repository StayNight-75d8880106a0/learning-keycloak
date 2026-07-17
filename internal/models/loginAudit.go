package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginAudit struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"`
	UserID    string    `json:"user_id" gorm:"type:char(36);not null;index"`
	Event     string    `json:"event" gorm:"type:varchar(30);not null"`
	IPAddress string    `json:"ip_address" gorm:"type:varchar(45)"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"createdAt"`
}

func (LoginAudit) TableName() string {
	return "login_audits"
}

func (u *LoginAudit) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
