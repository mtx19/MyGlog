package model

import (
	"MyGlog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Category Category `gorm:"foreignKey:ID;references:Cid;"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type: int not null" json:"cid"`
	Desc    string `gorm:"type: varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateAtr(data *Article) int {
	err = db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下所有文章
func GetArtsByCategory(cid int, pageSize, pageNum int) ([]Article, int, int64) {
	var data []Article
	offset := (pageNum - 1) * pageSize
	var total int64
	if pageSize == -1 {
		err = db.Preload("Category").Where("cid=?", cid).Find(&data).Count(&total).Error
	} else {
		err = db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid=?", cid).Find(&data).Count(&total).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		} else {
			return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
		}
	}
	return data, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return art, errmsg.ERROR_ART_NOT_EXIST
		} else {
			return art, errmsg.ERROR_ART_NOT_EXIST
		}
	}
	return art, errmsg.SUCCESS
}

// 查询文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int, int64) {
	var datas []Article
	offset := (pageNum - 1) * pageSize
	var total int64
	if pageSize == -1 {
		err = db.Preload("Category").Find(&datas).Count(&total).Error
	} else {
		err = db.Preload("Category").Limit(pageSize).Offset(offset).Find(&datas).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err:", err.Error())
		return nil, errmsg.ERROR, 0
	}
	return datas, errmsg.SUCCESS, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var c Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&c).Where("id=?", id).Updates(maps).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteArt(id int) int {
	var data Article
	err = db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
