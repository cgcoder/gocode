package models

import (
	"time"
)

type UserId int64

type User struct {
	ID            UserId
	Name          string
	Email         string
	Pass          string
	EmailVerified bool
	CreatedAt     time.Time
	PassResetAt   time.Time
}
