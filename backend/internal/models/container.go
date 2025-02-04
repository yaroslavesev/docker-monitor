package models

import (
	"time"
)

// Container struct
type Container struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IPAddress       string    `gorm:"not null" json:"ip_address"`
	LastPingTime    time.Time `json:"last_ping_time"`
	LastSuccessTime time.Time `json:"last_success_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
