package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByToken(ctx *gin.Context) {
	result, exists := ctx.Get("tokenInfo")
	if !exists {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "User not Found",
				"data":    nil,
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Get UserInfo Success",
			"data":    result.(map[string]interface{})["user"],
		},
	)
}
