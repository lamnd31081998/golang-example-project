package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	auth "golang-example-project/auth"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var master_db *gorm.DB

func init() {
	//env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file err === ", err)
	}

	//master db (postgres)
	dsn := os.Getenv("MASTER_DB_URL")
	master_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect db === ", err)
	}
}

func main() {
	router := gin.Default()

	auth.Init(router)

	router.Run(":" + os.Getenv("PORT"))
}
