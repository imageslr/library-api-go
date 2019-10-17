package image

import (
	"fmt"
	"library-api/config"
	userModel "library-api/models/user"
	"library-api/utils/file"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Upload 上传图片
func Upload(c *gin.Context) {
	image, err := c.FormFile("picture")
	if err != nil {
		fmt.Println("获取表单文件失败：", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "获取表单文件失败",
		})
		return
	}

	userInter, _ := c.Get("user")
	user := userInter.(userModel.User)

	fileName, ext := file.RandomFileName(image, strconv.Itoa(int(user.ID)), ".png")
	ext = strings.ToLower(ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".bmp" && ext != ".gif" {
		fmt.Println("文件格式错误，不能上传 " + ext + "格式的文件")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "文件格式错误，不能上传 " + ext + "格式的文件",
		})
		return
	}

	// 创建目录
	if _, err := os.Stat("storage/upload"); os.IsNotExist(err) {
		os.MkdirAll("storage/upload", os.ModePerm)
	}

	filePath := path.Join("storage/upload", fileName)

	// 上传文件到指定的 dst
	if err := c.SaveUploadedFile(image, filePath); err != nil {
		fmt.Println("保存文件失败", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "文件上传失败",
		})
		return
	}

	fmt.Println("保存文件到：", filePath)

	c.String(http.StatusOK, fmt.Sprintf("%s/upload/%s", config.AppConfig.URL, fileName))
}
