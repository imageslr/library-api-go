package user

import (
	"crypto/md5"
	"encoding/hex"
	"library-api/models"

	"github.com/Pallinder/go-randomdata"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name   string `json:"name" gorm:"type:varchar(255);not null" sql:"index"`
	Phone  string `json:"phone" gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Email  string `json:"email" gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Avatar string `json:"avatar" gorm:"type:varchar(255);not null"`
}

// BeforeCreate - hook
func (u *User) BeforeCreate() (err error) {
	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		u.Avatar = "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	}
	// 生成用户名
	if u.Name == "" {
		u.Name = randomdata.FullName(randomdata.RandomGender)
	}

	return err
}
