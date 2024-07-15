package routes

import (
	"golang-echo/handlers"
	"golang-echo/pkg/mysql"
	"golang-echo/repositories"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	Admin(e)
	Staff(e)
}

func H() *handlers.HandlerGlobal {
	repo := repositories.Repository(mysql.DB)

	handler := handlers.HandlerGlobal{
		AdminRepository: repo,
		StaffRepository: repo,
	}

	return handlers.Handler(handler)

}
