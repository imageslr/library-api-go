package user

import (
	"fmt"
	"library-api/config"
	"library-api/database"
	"library-api/models/user"
	"library-api/utils/cache"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	type Form struct {
		Phone string `form:"phone" binding:"required,numeric,len=11"`
	}

	var form Form
	if err := c.ShouldBindQuery(&form); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数非法",
		})
		return
	}

	if _, found := cache.Instance.Get(form.Phone); !found {
		rand.Seed(time.Now().UnixNano())
		code := rand.Intn(999999-100000) + 100000
		cache.Instance.SetDefault(form.Phone, code)
	}

	if code, found := cache.Instance.Get(form.Phone); found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint("验证码发送成功 ", code),
		})
	}
}

// Login - Check verify code and login or create a new user
func Login(c *gin.Context) {
	type Form struct {
		Phone string `form:"phone" binding:"required,numeric,len=11"`
		Code  string `form:"code" binding:"required,numeric,len=6"`
	}

	var form Form
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数非法",
		})
		return
	}

	code, found := cache.Instance.Get(form.Phone)
	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "验证码已过期",
		})
		return
	}
	if fmt.Sprintf("%v", code) != form.Code {
		fmt.Printf("code = %v, form.Code = %v", code, form.Code)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "验证码不正确",
		})
		return
	}

	cache.Instance.Delete(form.Phone)

	var u = user.User{Phone: form.Phone}
	var httpStatus = http.StatusOK
	if database.DB.Where("phone = ?", form.Phone).First(&u).RecordNotFound() {
		database.DB.Create(&u)
		httpStatus = http.StatusCreated
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": u.ID,
	})
	tokenString, err := token.SignedString([]byte(config.AppConfig.Key))
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(httpStatus, gin.H{
		"token": tokenString,
		"user":  u,
	})
}

// CurrentUser - 当前登录用户的信息
func CurrentUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
