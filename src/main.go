package main

import (
	"os"

	"github.com/gin-gonic/gin"

	authModule "golang-example-project/auth"
	configModule "golang-example-project/config"

	userModule "golang-example-project/user"
)

func init() {
	//Load ENV
	configModule.LoadEnv()

	//Connect DB
	configModule.ConnectMasterDB()

	//Migrate Table
	userModule.MigrateTable()
}

func main() {
	r := gin.Default()

	authModule.InitRouter(r)

	r.Run(":" + os.Getenv("PORT"))
}
