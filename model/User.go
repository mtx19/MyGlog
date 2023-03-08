package model

import (
	"MyGlog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type: varchar(20);not null" json:"password" validate:"required,min=6,max=16"`
	Nick     string `gorm:"type:varchar(20);not null" json:"nick"`
	Role     int    `gorm:"type: int;DEFAULT:2" json:"role" validate:"required,lte=2"` // 1 管理员  2 阅读者
	//Avatar   string //头像
}

// 查询用户是否存在
func CheckUserExist(name string) (int, int) {
	var users User
	db.Select("id").Where("username= ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED, int(users.ID)
	}
	return errmsg.SUCCESS, 0
}

// 新增
func CreateUser(data *User) int {
	scpwd := ScrypPwd(data.Password)
	if scpwd == "" {
		return errmsg.ERROR
	}
	data.Password = scpwd
	err = db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 {
		err = db.Find(&users).Error
	} else {
		err = db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err:", err.Error())
		return nil, 0
	}
	return users, total
}

// 编辑用户信息， 昵称，非修改密码  结构体更新只会更新非0值，所以要用map更新
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	maps["nick"] = data.Nick
	err = db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
const PwdKeyLen = 10

var salt []byte = []byte{1, 9, 8, 7, 6, 3, 0, 5}

func ScrypPwd(password string) string {
	hashpwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, PwdKeyLen)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(hashpwd)
	return fpw
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Password != ScrypPwd(password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
