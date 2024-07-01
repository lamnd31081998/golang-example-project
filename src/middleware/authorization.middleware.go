package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	repositoryModule "golang-example-project/repository"
	sharedModule "golang-example-project/shared"
	structModule "golang-example-project/struct"
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

	repositoryModule.UpdateUserById(map[string]interface{}{"ID": user.ID, "LastActive": time.Now(), "Status": 1})

	ctx.Set("tokenInfo", structModule.TokenInfo{Token: token, UserId: user.ID, User: *user})

	ctx.Next()
}
