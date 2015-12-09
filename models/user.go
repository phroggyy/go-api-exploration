package models

import "time"

type User struct {
    Id        int64   `json:"id"`
    Name      string  `json:"name"`
    Email     string  `json:"email"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

type Users []User