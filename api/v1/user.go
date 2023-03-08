package v1

import (
	"MyGlog/model"
	"MyGlog/utils/errmsg"
	"MyGlog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户
func AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	var msg string
	msg, code = validator.Validate(user)
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}
	code, _ = model.CheckUserExist(user.Username)
	if code == errmsg.SUCCESS {
		code = model.CreateUser(&user)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	var data model.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	code, existid := model.CheckUserExist(data.Username)
	if code == errmsg.SUCCESS || (existid == id && existid != 0) {
		code = model.EditUser(id, &data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func FindUserList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "-1"))
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
		"total":  total,
	})
}

// 删除
func DelUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
