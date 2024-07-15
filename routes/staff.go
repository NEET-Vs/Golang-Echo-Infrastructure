package routes

import (
	"github.com/labstack/echo/v4"
)

func Staff(e *echo.Group) {
	e.GET("/allstaff", H().AllStaff)
}
