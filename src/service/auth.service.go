package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
		ctx.JSON(
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

		ctx.JSON(
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
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Username duplicate",
				"data":    nil,
			},
		)

		return
	}

	var insertData structModule.User
	copier.Copy(&insertData, &requestBody)
	insertData.Password = _HashPassword(insertData.Password)
	user = repositoryModule.CreateUser(insertData)

	ctx.JSON(
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
		ctx.JSON(
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

		ctx.JSON(
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
		ctx.JSON(
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
		ctx.JSON(
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

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Create Token Fail",
				"data":    nil,
			},
		)

		return
	}

	user = repositoryModule.UpdateUserById(structModule.User{ID: user.ID, LastActive: time.Now()})

	ctx.JSON(
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

func _HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func _CheckHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
