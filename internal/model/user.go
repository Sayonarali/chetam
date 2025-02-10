package model

import "time"

type User struct {
	Id        int
	Login     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
