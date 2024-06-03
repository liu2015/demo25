package jztinfo

import (
	"demo25/global"
	"demo25/model/jztinfo"
)

type JztinfoService struct {
}

func (s *JztinfoService) Create(info jztinfo.Jztinfo) (err error) {
	err = global.DB.Create(&info).Error
	return err
}
