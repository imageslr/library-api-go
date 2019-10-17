package middlewares

import (
	"errors"
	"library-api/config"
	"library-api/database"
	"library-api/models/user"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth - 用户必须登录才能访问，否则返回 401
func Auth(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user", user)
	c.Next()
}

func getUser(c *gin.Context) (user.User, error) {
	var user user.User
	var err = errors.New("未登录")

	tokenString := c.GetHeader(config.AppConfig.TokenKey)
	if tokenString == "" {
		return user, err
	}

	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.Key), nil
	})
	if tokenErr != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		if database.DB.Where("id = ?", userID).First(&user).RecordNotFound() {
			return user, errors.New("内部错误")
		}
		return user, nil
	}

	return user, err
}
