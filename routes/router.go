package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DelCategory)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DelArticle)
		// 文件上传
		auth.POST("upload", v1.UpLoad)

	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("categorys", v1.GetCategorys)
		router.GET("article/list/:id", v1.GetCateArticle)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.GET("articles", v1.GetArticles)
		// 登录模块的路由接口
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
