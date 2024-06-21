package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	configModule "golang-example-project/config"
	repositoryModule "golang-example-project/repository"
	routerModule "golang-example-project/router"
)

func init() {
	//Load ENV
	configModule.LoadEnv()

	//Connect Redis
	configModule.ConnectRedis()

	//Connect DB
	configModule.ConnectMasterDB()

	//Migrate Table
	repositoryModule.MigrateUserTable()
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	routerModule.InitAuthRouter(r)
	routerModule.InitUserRouter(r)

	r.Run(":" + os.Getenv("PORT"))
}
