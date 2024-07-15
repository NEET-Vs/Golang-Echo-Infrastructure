package database

import (
	"fmt"
	"golang-echo/models"
	"golang-echo/pkg/mysql"
)

func RunMigration() {

	var err error
	err = mysql.DB.AutoMigrate(
		&models.Admin{},
		&models.Staff{},
	)

	if err != nil {
		fmt.Println(err)
		panic("DB Migration Failed ")
	}

	fmt.Println("All Migration Success")
}
