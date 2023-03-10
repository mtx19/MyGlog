package v1

import (
	"MyGlog/middleware"
	"MyGlog/model"
	"MyGlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var data model.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		errmsg.SendBinJsonError(ctx)
		return
	}
	var token string
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}
