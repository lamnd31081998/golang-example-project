package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	sf "github.com/sa-/slicefunk"
	"golang.org/x/crypto/bcrypt"

	repositoryModule "golang-example-project/repository"
	structModule "golang-example-project/struct"

	"github.com/jinzhu/copier"

	sharedModule "golang-example-project/shared"
)

func Register(ctx *gin.Context) {
	var requestBody structModule.RegisterDto
	if err := ctx.BindJSON(&requestBody); err != nil {
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

	user := repositoryModule.FindUserByUsername(requestBody.Username)
	if user != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Username duplicate",
				"data":    nil,
			},
		)

		return
	}

	if requestBody.Password != requestBody.Confirm_Password {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Password and Confirm Password unmatch",
				"data":    nil,
			},
		)

		return
	}

	var insertData structModule.User
	copier.Copy(&insertData, &requestBody)
	insertData.Password = _HashPassword(insertData.Password)
	user = repositoryModule.CreateUser(insertData)

	ctx.AbortWithStatusJSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "User Created",
			"data":    user,
		},
	)
}

func Login(ctx *gin.Context) {
	var requestBody structModule.LoginDto
	if err := ctx.BindJSON(&requestBody); err != nil {
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

	user := repositoryModule.FindUserByUsername(requestBody.Username)
	if user == nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Cannot Find UserInfo",
				"data":    nil,
			},
		)

		return
	}

	if !_CheckHashPassword(requestBody.Password, user.Password) {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Wrong Password",
				"data":    nil,
			},
		)

		return
	}

	token, err := sharedModule.JwtCreateToken(user.ID, user.Username)
	if err != nil {
		fmt.Println("Create Token Fail === ", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Create Token Fail",
				"data":    nil,
			},
		)

		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Login Success",
			"data": gin.H{
				"user":         user,
				"access_token": token,
			},
		},
	)
}

func LogoutByToken(ctx *gin.Context) {
	result, _ := ctx.Get("tokenInfo")
	result_parsed := result.(structModule.TokenInfo)
	user_info := result_parsed.User

	repositoryModule.UpdateUserById(structModule.User{ID: user_info.ID, Status: 2})

	if err := sharedModule.DelRedisByKey(result_parsed.Token); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Logout Fail",
				"data":    nil,
			},
		)

		return
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Logout Success",
			"data":    nil,
		},
	)
}

func _HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func _CheckHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
