package router

import (
	"github.com/gin-gonic/gin"

	middlewareModule "golang-example-project/middleware"
	serviceModule "golang-example-project/service"
)

func InitUserRouter(route *gin.Engine) {
	api := route.Group("/user")
	{
		api.GET("", middlewareModule.CheckAuthorization, serviceModule.GetUserByToken)
		api.PUT("", middlewareModule.CheckAuthorization, serviceModule.UpdateUserByToken)
	}
}
