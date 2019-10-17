package file

import (
	"library-api/utils"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

// RandomFileName 生成随机文件名
// filePrefix 文件名前缀
// defaultExt 默认 ext (有些文件上传时没带文件后缀)
func RandomFileName(f *multipart.FileHeader, filePrefix, defaultExt string) (string, string) {
	fileName := f.Filename
	// 获取文件的后缀名，因图片从剪贴板里黏贴时后缀名为空，所以此处确保后缀一直存在
	ext := path.Ext(fileName)
	if ext == "" && defaultExt != "" {
		ext = defaultExt // 默认的后缀
	}

	// 拼接文件名，加前缀是为了增加辨析度，前缀可以是相关数据模型的 ID
	// 格式：Prefix_Timestamp_RandString.Ext
	// 值如：1_1493521050_7BVc9v9ujP.png
	fileName = filePrefix + "_" +
		strconv.Itoa(int(time.Now().Unix())) + "_" +
		string(utils.RandomBytes(10)) + ext

	return fileName, ext
}
