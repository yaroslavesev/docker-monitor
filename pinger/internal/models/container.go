package models

import (
	"time"
)

// Container struct
type Container struct {
	ID              uint      `json:"id"`
	IPAddress       string    `json:"ip_address"`
	LastPingTime    time.Time `json:"last_ping_time"`
	LastSuccessTime time.Time `json:"last_success_time"`
}
