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
	Nickname  string    `json:"nickname" gorm:"type:varchar(255);not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique;default:NULL" sql:"index"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(255);not null"`
	Openid    string    `json:"openid" gorm:"type:varchar(50);not null"`
	Address   string    `json:"address" gorm:"type:varchar(255);not null"`
	Birthday  string    `json:"birthday" gorm:"type:varchar(20);default:Null"`
	IDNumber  string    `json:"id_number" gorm:"type:varchar(20);not null"`
	IDCardImg IDCardImg `json:"id_card_img" gorm:"type:varchar(255)"`
	Postcode  string    `json:"postcode" gorm:"type:varchar(10);not null"`
}

// BeforeCreate -
func (u *User) BeforeCreate() (err error) {
	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		u.Avatar = "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	}
	// 生成用户名
	if u.Nickname == "" {
		u.Nickname = randomdata.FullName(randomdata.RandomGender)
	}

	return err
}
