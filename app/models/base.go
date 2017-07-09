package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

var Database *gorm.DB

func InitDB () {
	fmt.Println("application started")
	var err error

	Database, err = gorm.Open("postgres", "host=localhost user=nikitagromov dbname=gorm sslmode=disable password=")

	if err != nil {
		panic("failed to connect database")
	}

	Database.DB().Ping()
	//Database.CreateTable(&models.Task{})
	Database.AutoMigrate(&Task{})
	Database.AutoMigrate(&Project{})
	Database.AutoMigrate(&User{})

}

