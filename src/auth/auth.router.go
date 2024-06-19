package auth

import "github.com/gin-gonic/gin"

func InitRouter(route *gin.Engine) {
	api := route.Group("/")
	{
		api.POST("/register", Register)
		api.POST("/login")
	}
}
