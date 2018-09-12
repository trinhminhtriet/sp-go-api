package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/trinhminhtriet/sp-go-api/configs"
)

func main() {
	// Get local env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init db
	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.AutoMigrate(
	//	&models.User{},
	//	&models.Article{},
	//	&models.Tag{})

	// Init router
	router := configs.BuildRoutes(db.Debug())
	router.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
