package models

import "time"

type Config struct {
	BackendURL   string
	PingInterval time.Duration
}
