package shared

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtCreateToken(user_id uint, username string) (string, error) {
	expire_in_day, err := strconv.Atoi(os.Getenv("EXPIRE_IN"))
	if err != nil {
		return "", err
	}
	duration := expire_in_day * 24 * 60 * 60 * 1000000000

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user_id,
		"username":   username,
		"expired_at": time.Now().Add(time.Duration(duration)),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	payload, err := json.Marshal(gin.H{
		"user_id":  user_id,
		"username": username,
	})
	if err != nil {
		return "", err
	}

	if err := SetRedisByKey(tokenString, payload, duration); err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtParseToken(token string) (map[string]interface{}, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return result.Claims.(jwt.MapClaims), nil
}
