package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	repositoryModule "golang-example-project/repository"
	sharedModule "golang-example-project/shared"
)

func CheckAuthorization(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Missing Token",
				"data":    nil,
			},
		)

		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)

	jwt_parsed, err := sharedModule.JwtParseToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Token Invalid",
				"data":    nil,
			},
		)

		return
	}

	cache_data, err := sharedModule.GetRedisByKey(token)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Token Expired",
				"data":    nil,
			},
		)

		return
	}

	var cache_parsed map[string]interface{}
	json.Unmarshal([]byte(cache_data), &cache_parsed)

	if jwt_parsed["user_id"] != cache_parsed["user_id"] || jwt_parsed["username"] != cache_parsed["username"] {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Token Unmatch",
				"data":    nil,
			},
		)

		return
	}

	user := repositoryModule.FindUserById(uint(cache_parsed["user_id"].(float64)))
	if user == nil || user.Username != cache_parsed["username"] {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "User not Found",
				"data":    nil,
			},
		)

		return
	}
	ctx.Set("tokenInfo", map[string]interface{}{"user_id": user.ID, "user": user, "token": token})

	ctx.Next()
}
