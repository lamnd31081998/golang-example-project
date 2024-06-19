package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	sf "github.com/sa-/slicefunk"
)

func Register(c *gin.Context) {
	var requestBody RegisterDto
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Get Data Fail",
				"data":    nil,
			},
		)

		return
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := sf.Map(validationErrors, func(item validator.FieldError) ValidateError {
			return ValidateError{
				Field: strings.ToLower(item.Field()),
				Type:  strings.ToLower(item.ActualTag()),
			}
		})

		c.JSON(
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

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "OK",
			"data":    nil,
		},
	)
}
