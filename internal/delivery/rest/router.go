package rest

import "github.com/labstack/echo"

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/menu", handler.GetMenu)
}