package router

import (
	"github.com/gin-gonic/gin"

	serviceModule "golang-example-project/service"
)

func InitAuthRouter(route *gin.Engine) {
	api := route.Group("/")
	{
		api.POST("/register", serviceModule.Register)
		api.POST("/login", serviceModule.Login)
	}
}
