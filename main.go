package main

import (
	"btpn-final/config"
	"btpn-final/models"
	"btpn-final/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	r := router.SetupRouter(db)
	r.Run()
}
