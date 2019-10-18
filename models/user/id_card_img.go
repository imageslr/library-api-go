package user

import (
	"database/sql/driver"
	"encoding/json"
)

// IDCardImg 身份证照片 url
type IDCardImg struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

// Scan 实现 gorm Scanner 接口
func (ls *IDCardImg) Scan(value interface{}) error {
	if value == nil {
		// *ls = IDCardImg{}
		return nil
	}
	t := IDCardImg{}
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	*ls = t
	return nil
}

// Value 实现 gorm Valuer 接口
func (ls IDCardImg) Value() (driver.Value, error) {
	b, e := json.Marshal(ls)
	return b, e
}
