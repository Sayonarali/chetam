package handlers

import (
	"chetam/internal/apperror"
	"chetam/internal/model"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

type RoutesInterface interface {
}

func GetRoutesList(logger *slog.Logger, routes RoutesInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.GetRoutesList"

		var list []model.Route

		route1 := model.Route{
			Name:     "Необычные места России",
			Rate:     5.0,
			Author:   "Лизочка",
			Duration: 120,
		}
		route2 := model.Route{
			Name:     "Природные чудеса",
			Rate:     4.7,
			Author:   "Лизочка",
			Duration: 10,
		}
		route3 := model.Route{
			Name:     "Исторические памятники",
			Rate:     3.9,
			Author:   "Лизочка",
			Duration: 40,
		}
		list = append(list, route1, route2, route3)

		return c.JSON(http.StatusOK, list)
	}
}

func GetRoute(logger *slog.Logger, routes RoutesInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.GetRoute"

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			e := model.Error{
				Error: apperror.ErrBadRequest,
			}

			logger.Error("failed to decode request body",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusBadRequest, e)
		}
		var route model.Route

		switch id {
		case 1:
			point1 := model.Point{
				Name: "Гигантский смартфон в Омске",
				Lat:  55.015516,
				Long: 73.425020,
			}
			point2 := model.Point{
				Name: "Лесная надпись «Ленину 100 лет» в Курганской области",
				Lat:  54.28632,
				Long: 64.474820,
			}
			point3 := model.Point{
				Name: "А это город Остров",
				Lat:  57.34364,
				Long: 28.35743,
			}

			route = model.Route{
				Name:   "Необычные места России",
				Rate:   5.0,
				Points: []model.Point{point1, point2, point3},
			}

		case 2:
			point4 := model.Point{
				Name: "Кижи (Музей-заповедник деревянного зодчества)",
				Lat:  62.066876,
				Long: 35.225357,
			}
			point5 := model.Point{
				Name: "Долина гейзеров (Камчатка)",
				Lat:  54.430663,
				Long: 160.137676,
			}
			point8 := model.Point{
				Name: "Кунгурская ледяная пещера",
				Lat:  57.440000,
				Long: 57.006944,
			}
			point9 := model.Point{
				Name: "Столбы выветривания (Плато Мань-Пупу-нёр)",
				Lat:  62.257222,
				Long: 59.299167,
			}
			point10 := model.Point{
				Name: "Эльбрус (Высшая точка России)",
				Lat:  43.355075,
				Long: 42.439716,
			}

			route = model.Route{
				Name:   "Природные чудеса",
				Rate:   5.0,
				Points: []model.Point{point4, point5, point8, point9, point10},
			}
		case 3:
			point6 := model.Point{
				Name: "Медный всадник (Санкт-Петербург)",
				Lat:  59.936389,
				Long: 30.302222,
			}
			point7 := model.Point{
				Name: "Мамаев курган и Родина-мать (Волгоград)",
				Lat:  48.742183,
				Long: 44.537151,
			}

			route = model.Route{
				Name:   "Исторические памятники",
				Rate:   5.0,
				Points: []model.Point{point6, point7},
			}
		}

		return c.JSON(http.StatusOK, route)
	}
}
