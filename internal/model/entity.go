package model

import "time"

// User - модель пользователя
type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Route - модель маршрута
type Route struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	AuthorId      uint      `json:"creatorId"` // id User пользователя-создателя
	CityId        uint      `json:"cityId"`
	Points        []Point   `json:"points"`        // Массив id Point в порядке маршрута
	TotalTime     int       `json:"totalTime"`     // Общее время в минутах
	IsPublic      bool      `json:"isPublic"`      // Доступен ли для других пользователей
	Filters       string    `json:"filters"`       // JSON-строка с фильтрами (для генерации)
	AverageRating float64   `json:"averageRating"` // Вычисляемое поле (не в БД)
	CreatedAt     time.Time `json:"createdAat"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// Point - точка интереса (достопримечательность)
type Point struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Latitude         float64   `json:"latitude"` // Геокоординаты
	Longitude        float64   `json:"longitude"`
	CityId           uint      `json:"cityId"`
	CategoryId       uint      `json:"category"`         // e.g., "museum", "park", "restaurant"
	AverageVisitTime int       `json:"averageVisitTime"` // В минутах
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// Category - категория точек e.g., "museum", "park", "restaurant"
type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

// City - модель города
type City struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

// Rating - модель оценки маршрута
type Rating struct {
	Id        uint      `json:"id"`
	RouteId   uint      `json:"routeId"`
	UserId    uint      `json:"userId"`
	Score     int       `json:"score"` // 1-5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Error struct {
	Error string `json:"error"`
}
