package jztinfo

import (
	"demo25/global"
	"demo25/model/jztinfo"
	"demo25/model/jztinfo/request"
)

type JztinfoService struct {
}

func (s *JztinfoService) Create(info jztinfo.Jztinfo) (err error) {
	err = global.DB.Create(&info).Error
	return err
}

func (s *JztinfoService) Delete(id string) (err error) {
	err = global.DB.Where("id =? ", id).Delete(&jztinfo.Jztinfo{}).Error
	return err
}

func (s *JztinfoService) Deletes(ids []string) (err error) {
	err = global.DB.Where("id in ?", ids).Delete(&[]jztinfo.Jztinfo{}).Error
	return err
}
func (s *JztinfoService) Updates(info jztinfo.Jztinfo) (err error) {
	err = global.DB.Where("id =?", info.ID).Updates(&info).Error
	return err
}

func (s *JztinfoService) FindOne(id string) (list jztinfo.Jztinfo, err error) {
	err = global.DB.Where("id =?", id).First(&list).Error
	return list, err
}

func (s *JztinfoService) FindList(info request.Findjztinfo) (list []jztinfo.Jztinfo, total int64, err error) {
	limts := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(jztinfo.Jztinfo{})
	var infos []jztinfo.Jztinfo

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name like ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limts != 0 {
		db = db.Limit(limts).Offset(offset)
	}
	err = db.Find(&infos).Error
	return infos, total, err
}
