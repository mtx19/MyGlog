package v1

import (
	"MyGlog/model"
	"MyGlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileheader, _ := c.Request.FormFile("file")
	filesize := fileheader.Size
	url, code := model.UploadFile(file, filesize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   url,
	})
}
