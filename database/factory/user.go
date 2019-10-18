package factory

import (
	"fmt"

	"library-api/database"
	"library-api/models/user"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	// 头像假数据
	avatars = []string{
		"http://qiniu.library-online.cn/avatar/1.png",
		"http://qiniu.library-online.cn/avatar/2.png",
		"http://qiniu.library-online.cn/avatar/3.png",
		"http://qiniu.library-online.cn/avatar/4.png",
		"http://qiniu.library-online.cn/avatar/5.png",
		"http://qiniu.library-online.cn/avatar/6.png",
	}
)

func userFactory(i int) *factory.Factory {
	u := &user.User{}

	return factory.NewFactory(
		u,
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return fmt.Sprintf("user-%d", i+1), nil
	}).Attr("Avatar", func(args factory.Args) (interface{}, error) {
		return avatars[randomdata.Number(0, len(avatars))], nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		return randomdata.Email(), nil
	}).Attr("Phone", func(args factory.Args) (interface{}, error) {
		if i == 1 {
			return "13000000000", nil
		}
		return randomdata.PhoneNumber(), nil
	}).Attr("IDCardImg", func(args factory.Args) (interface{}, error) {
		if i%2 == 0 {
			return user.IDCardImg{
				Front: randomdata.IpV4Address(),
				Back:  randomdata.IpV4Address(),
			}, nil
		} else {
			return user.IDCardImg{}, nil
		}
	})
}

// UsersTableSeeder -
func UsersTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&user.User{})
	}

	for i := 0; i < 10; i++ {
		user := userFactory(i).MustCreate().(*user.User)
		if err := database.DB.Create(&user).Error; err != nil {
			fmt.Printf("用户创建失败: %v", err)
		}
	}

}
