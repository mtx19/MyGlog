package routers

import (
	v1 "MyGlog/api/v1"
	"MyGlog/middleware"
	"MyGlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//用户模块
		Auth.POST("user/add", v1.AddUser)
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DelUser)
		//分类模块
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DelCategory)
		//文章模块
		Auth.POST("art/add", v1.AddArt)
		Auth.DELETE("art/:id", v1.DelArt)
		Auth.PUT("art/:id", v1.EditArt)
		//上传文件
		Auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.POST("login", v1.Login)
		router.GET("users/", v1.FindUserList)
		router.GET("categorys/", v1.FindCategoryList)
		router.GET("arts/", v1.FindArtList)
		router.GET("artinfo/:id", v1.GetArtInfo)
		router.GET("art/catelist", v1.GetArtsByCategory)
	}

	r.Run(utils.HttpPort)
}
