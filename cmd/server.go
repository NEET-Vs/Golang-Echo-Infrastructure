package cmd

import (
	"golang-echo/database"
	"golang-echo/pkg/mysql"
	"golang-echo/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()

	mysql.DataBaseinit()
	database.RunMigration()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	routes.RouteInit(e.Group(""))
	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "5000"
	}

	e.Logger.Fatal(e.Start(":" + PORT))
}
