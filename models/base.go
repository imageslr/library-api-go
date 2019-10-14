package models

import "github.com/jinzhu/gorm"

// BaseModel model 基类
type BaseModel struct {
	gorm.Model
}
