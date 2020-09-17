package router

import (
	"LearnLogin/model"
	"LearnLogin/utils"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	//设置跨域中间件
	router := r.Group("/v1", utils.Cors())
	{
		router.POST("/login", model.Login)
		router.GET("/", model.Home)
	}

	//需要token的路由组
	captcha := r.Group("/v2", utils.Cors(), utils.JwtAuth())
	{
		//添加用户
		captcha.POST("/registered", model.Registered)
		//删除用户
		captcha.DELETE("/delete", model.Delete)
		//更新用户信息
		captcha.PUT("/updateuserinfo", model.UpdateUserInfo)
		//遍历所有用户信息
		captcha.POST("/alluserinfo", model.AllUserInfo)
		//创建分类

		//更新分类

		//删除分类

		//遍历分类

		//创建题目

		//更新题目

		//删除题目

	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
