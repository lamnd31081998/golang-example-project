package service

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	sf "github.com/sa-/slicefunk"

	"golang-example-project/repository"
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

func UpdateUserByToken(ctx *gin.Context) {
	var requestBody structModule.UpdateUserInfoByTokenDto
	if err := ctx.ShouldBindWith(&requestBody, binding.FormMultipart); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Convert Data Fail",
				"data":    nil,
			},
		)

		return
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := sf.Map(validationErrors, func(item validator.FieldError) gin.H {
			return gin.H{
				"field": strings.ToLower(item.Field()),
				"type":  strings.ToLower(item.ActualTag()),
			}
		})

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Validate Data Fail",
				"data":    nil,
				"errors":  errors,
			},
		)

		return
	}

	tokenInfo_raw, _ := ctx.Get("tokenInfo")
	user_info := tokenInfo_raw.(structModule.TokenInfo).User
	user_UpdateData := map[string]interface{}{"ID": user_info.ID, "Name": requestBody.Name, "AvatarUrl": ""}

	if requestBody.File != nil {
		filename := strconv.Itoa(int(time.Now().Unix())) + "_" + requestBody.File.Filename
		if err := ctx.SaveUploadedFile(requestBody.File, "./../public/avatar/"+filename); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": "Upload File Err",
					"data":    nil,
				},
			)

			return
		}

		AvatarUrl := os.Getenv("SERVICE_URL") + "/assets/avatar/" + filename
		user_UpdateData["AvatarUrl"] = AvatarUrl
	}

	if requestBody.File == nil && requestBody.AvatarUrl == "" {
		user_UpdateData["AvatarUrl"] = nil
	}

	repository.UpdateUserById(user_UpdateData)

	user := repository.FindUserById(user_info.ID)
	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update UserInfo Success",
			"data":    user,
		},
	)
}
