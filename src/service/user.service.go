package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	structModule "golang-example-project/struct"
)

func GetUserByToken(ctx *gin.Context) {
	result, _ := ctx.Get("tokenInfo")
	result_parsed := result.(structModule.TokenInfo)

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Get UserInfo Success",
			"data":    result_parsed.User,
		},
	)
}
