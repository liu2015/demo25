package jztinfo

import "gorm.io/gorm"

type Jztinfo struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Age  string `json:"age" form:"age" gorm:"column:age;comment:;"`
	Make string `json:"make" form:"make" gorm:"column:make;comment:;"`
	Info string `json:"info" form:"info" gorm:"column:info;comment:;"`
}

func (Jztinfo) TableName() string {
	return "jztinfos"
}
