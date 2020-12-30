package model

import (
    "time"
)

type UserId int64

type User struct {
    Id UserId
    Name string
    Email string
    Pass string
    EmailVerified bool
    CreatedAt Time
    PassResetAt Time
}
