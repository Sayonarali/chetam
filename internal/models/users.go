package models

import "time"

type User struct {
	Id        uint
	Login     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
