package v1

import (
	"MyGlog/model"
	"MyGlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArt(ctx *gin.Context) {
	var data model.Article
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	code = model.CreateAtr(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
	})
}

// 编辑文章
func EditArt(ctx *gin.Context) {
	var data model.Article
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.PARAMERROR,
			"msg":    errmsg.GetErrMsg(errmsg.PARAMERROR),
		})
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.EditArt(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询分类下的文章
func GetArtsByCategory(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Query("cid"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "-1"))
	data, code, total := model.GetArtsByCategory(cid, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
		"total":  total,
	})
}

// 查询单个文章
func GetArtInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArtInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
	})
}

// 查询文章列表
func FindArtList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagenum", "-1"))
	data, code, total := model.GetArts(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
		"total":  total,
	})
}

// 删除
func DelArt(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
