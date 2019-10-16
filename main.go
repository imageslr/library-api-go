package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"

	"library-api/config"
	"library-api/database"
	"library-api/database/factory"
	"library-api/models/user"
	"library-api/routes"
)

var (
	// 生成 mock 数据，只在非 release 时生效
	needMock = pflag.BoolP("mock", "m", false, "need mock data")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 初始化配置
	config.InitConfig("", true)

	// 设置 gin
	g := gin.New()
	setupGin(g)

	// 设置路由
	routes.Register(g)

	// 设置数据库
	db, err := setupDB()
	if err != nil {
		return
	}
	defer db.Close()

	if *needMock {
		// 生成 mock 数据
		if config.AppConfig.RunMode == config.RunmodeRelease {
			panic("[mock] 请在非生产环境中使用")
		}
		fmt.Print("\n\n生成 mock 数据中...\n\n")
		factory.Mock()
		fmt.Print("\n\n成功生成 mock 数据！\n\n")
	} else {
		// 启动服务器
		fmt.Printf("\n\n Start to listening the incoming requests on http address: %s n\n", config.AppConfig.Addr)
		if err := http.ListenAndServe(config.AppConfig.Addr, g); err != nil {
			log.Fatal("http server 启动失败", err)
		}
	}

}

func setupGin(g *gin.Engine) {
	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)
}

func setupDB() (*gorm.DB, error) {
	db := database.InitDB()

	// db migrate
	db.AutoMigrate(
		&user.User{},
	)

	return db, nil
}
