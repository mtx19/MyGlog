package middleware

import (
	"MyGlog/utils"
	"MyGlog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyCliams struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetCliams := MyCliams{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "superM",
		},
	}
	reqCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, SetCliams)
	token, err := reqCliams.SignedString(JwtKey)
	if err != nil {
		log.Println(err.Error())
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS

}

// 验证token
func CheckToken(token string) (*MyCliams, int) {
	settoken, err := jwt.ParseWithClaims(token, &MyCliams{}, func(*jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, errmsg.ERROR
	}
	if settoken != nil {
		if key, ok := settoken.Claims.(*MyCliams); ok && settoken.Valid {
			return key, errmsg.SUCCESS
		}
	}
	return nil, errmsg.ERROR
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code = errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_FORMAT_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, checkcode := CheckToken(checkToken[1])
		if checkcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME_OUT
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
