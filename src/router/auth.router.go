package router

import (
	"github.com/gin-gonic/gin"

	middlewareModule "golang-example-project/middleware"
	serviceModule "golang-example-project/service"
)

func InitAuthRouter(route *gin.Engine) {
	api := route.Group("/")
	{
		api.POST("/register", serviceModule.Register)
		api.POST("/login", serviceModule.Login)
		api.DELETE("/logout", middlewareModule.CheckAuthorization, serviceModule.LogoutByToken)
	}
}
