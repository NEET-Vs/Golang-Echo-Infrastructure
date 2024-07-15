package routes

import (
	"github.com/labstack/echo/v4"
)

func Admin(e *echo.Group) {

	e.GET("/alladmin", H().AllAdmin)
}
