package user

import (
	"fmt"
	"library-api/utils/cache"
	"math/rand"
	"net/http"
	"time"

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
