package mysql

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DataBaseinit() {
	var err error

	// env database relation
	var dbconfig = os.Getenv("DB_CHARNTIME")
	var mst = os.Getenv("DB_MST")

	// env DATABASE HOST
	var dbuser = os.Getenv("DB_USER")
	var dbpass = os.Getenv("DB_PASS")
	var dbhost = os.Getenv("DB_HOST")
	var dbport = os.Getenv("DB_PORT")
	var LocationTime = os.Getenv("LOC_TIME")

	loc := url.QueryEscape(LocationTime)
	if err != nil {
		panic(err)
	}
	var mysqlconfig = dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/"
	var gormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s%s%s%s", mysqlconfig, mst, dbconfig, loc)), gormConfig)
	if err != nil {
		fmt.Println("connect to database mst failed")
		DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:password@tcp(127.0.0.1:3306)/mst_test?charset=utf8mb4&parseTime=True&loc=Asia/Jakarta")), gormConfig)
	}
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("connected to database")
}
