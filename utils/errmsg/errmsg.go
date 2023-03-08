package errmsg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS    = 200
	ERROR      = 500
	PARAMERROR = 501
	//code = 1000... 用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_RUNTIME_OUT  = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_FORMAT_WRONG = 1007
	ERROR_USER_NO_RIGHT      = 1008

	//code = 2000... 文章模块错误
	ERROR_ART_NOT_EXIST = 2001

	//code = 3000... 分类模块错误
	ERROR_CATEGORY_USED      = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "Fail",
	PARAMERROR:               "Param bind error",
	ERROR_USERNAME_USED:      "用户名已存在！",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_TOKEN_NOT_EXIST:    "Token 不存在",
	ERROR_TOKEN_RUNTIME_OUT:  "登录校验过期",
	ERROR_TOKEN_WRONG:        "登录校验错误",
	ERROR_TOKEN_FORMAT_WRONG: "TOKEN 格式错误",
	ERROR_CATEGORY_USED:      "分类已经存在",
	ERROR_CATEGORY_NOT_EXIST: "文章分类不存在",
	ERROR_ART_NOT_EXIST:      "文章不存在",
	ERROR_USER_NO_RIGHT:      "用户权限不足",
}

func GetErrMsg(code int) string {
	if res, ok := codeMsg[code]; ok {
		return res
	}
	return fmt.Sprintf("未知错误，错误码：%d", code)
}
func SendBinJsonError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": PARAMERROR,
		"msg":  GetErrMsg(PARAMERROR),
	})
}
