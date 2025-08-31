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

type Route struct {
	Name     string
	Rate     float32
	Points   []Point
	Author   string
	Duration time.Duration
}

type Point struct {
	Name string
	Lat  float32
	Long float32
}

type Error struct {
	Error string `json:"error"`
}
