package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	sf "github.com/sa-/slicefunk"
	"golang.org/x/crypto/bcrypt"

	userModule "golang-example-project/user"

	"github.com/jinzhu/copier"
)

func Register(ctx *gin.Context) {
	var requestBody RegisterDto
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

	user := userModule.FindByUsername(requestBody.Username)
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

	var insertData userModule.User
	copier.Copy(&insertData, &requestBody)
	insertData.Password = _HashPassword(insertData.Password)
	user = userModule.Create(insertData)

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "User Created",
			"data":    user,
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
