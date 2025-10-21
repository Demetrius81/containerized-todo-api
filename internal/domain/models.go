package domain

import (
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Done        bool      `gorm:"not null;default:false" json:"done"`
	CreatedAt   time.Time `gorm:"not null;default:now();autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null;default:now();autoUpdateTime" json:"updated_at,omitempty"`
}
