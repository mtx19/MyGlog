package model

import (
	"MyGlog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20) not null" json:"name"`
}

// 查询分类是否存在
func CheckCategoryExist(name string) (int, int) {
	var data Category
	db.Select("id").Where("name= ?", name).First(&data)
	if data.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED, int(data.ID)
	}
	return errmsg.SUCCESS, 0
}

// 新增
func CreateCategory(data *Category) int {
	err = db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategorys(pageSize int, pageNum int) ([]Category, int64) {
	var datas []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 {
		err = db.Find(&datas).Error
	} else {
		err = db.Limit(pageSize).Offset(offset).Find(&datas).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err:", err.Error())
		return nil, 0
	}
	return datas, total
}

// 编辑分类信息， 昵称，非修改密码  结构体更新只会更新非0值，所以要用map更新
func EditCategory(id int, data *Category) int {
	var c Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&c).Where("id=?", id).Updates(maps).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteCategory(id int) int {
	var data Category
	err = db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
