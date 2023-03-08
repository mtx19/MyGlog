package v1

import (
	"MyGlog/model"
	"MyGlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户
func AddCategory(ctx *gin.Context) {
	var data model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	code, _ = model.CheckCategoryExist(data.Name)
	if code == errmsg.SUCCESS {
		code = model.CreateCategory(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
	})
}

// 编辑用户
func EditCategory(ctx *gin.Context) {
	var data model.Category
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	code, existid := model.CheckUserExist(data.Name)
	if code == errmsg.SUCCESS || (existid == id && existid != 0) {
		code = model.EditCategory(id, &data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询分类
func FindCategoryList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "-1"))
	data, total := model.GetCategorys(pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
		"total":  total,
	})
}

// 删除
func DelCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
