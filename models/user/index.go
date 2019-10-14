package user

import (
	"crypto/md5"
	"encoding/hex"
	"library-api/models"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name   string `gorm:"type:varchar(255);not null" sql:"index"`
	Phone  string `gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Email  string `gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Avatar string `gorm:"type:varchar(255);not null"`
}

// BeforeCreate - hook
func (u *User) BeforeCreate() (err error) {
	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		u.Avatar = "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	}

	return err
}
