package user

import (
	"fmt"
	"library-api/config"
	"library-api/database"
	userModel "library-api/models/user"
	"library-api/utils/cache"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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

// Login checks verify code and login or creates a new user
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

	var u = userModel.User{Phone: form.Phone}
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

// CurrentUser 当前登录用户的信息
func CurrentUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

// UpdateCurrentUser 更新用户信息
func UpdateCurrentUser(c *gin.Context) {
	type Form struct {
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Name      string `json:"name"`
		Birthday  string `json:"birthday"`
		IDNumber  string `json:"id_number"`
		Address   string `json:"address"`
		Postcode  string `json:"postcode"`
		IDCardImg struct {
			Front string `json:"front"`
			Back  string `json:"back"`
		} `json:"id_card_img"`
	}

	userInter, _ := c.Get("user")
	user := userInter.(userModel.User)

	var form Form
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数非法",
		})
		return
	}

	if form.Nickname != "" {
		user.Nickname = form.Nickname
	}
	if form.Avatar != "" {
		user.Avatar = form.Avatar
	}
	if form.Name != "" {
		user.Name = form.Name
	}
	if form.Birthday != "" {
		user.Birthday = form.Birthday
	}
	if form.IDNumber != "" {
		user.IDNumber = form.IDNumber
	}
	if form.Address != "" {
		user.Address = form.Address
	}
	if form.Postcode != "" {
		user.Postcode = form.Postcode
	}
	if form.IDCardImg.Front != "" && form.IDCardImg.Back != "" {
		user.IDCardImg = userModel.IDCardImg(form.IDCardImg)
	}

	if err := database.DB.Model(&user).Update(&user).Error; err != nil {
		log.Warnf("用户更新失败: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "用户更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
