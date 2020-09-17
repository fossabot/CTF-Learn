package model

import (
	"LearnLogin/config"
	"LearnLogin/database"
	"LearnLogin/service"
	"LearnLogin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册用户
func Registered(c *gin.Context) {
	var data database.User
	var message string
	//使用bindjson填充数据
	c.BindJSON(&data)
	//验证用户名是否存在
	userCode := database.CheckUser(data.Username, data.Role)
	if userCode == config.Success {
		//创建用户
		service.AddUser(&data)
		message = "注册成功"
	}
	if userCode == config.UserExisted {
		c.Abort()
		message = "用户已存在"
	}
	if userCode == config.NoRole {
		c.Abort()
		message = "权限错误"
	}
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":    userCode,
		"message": message,
	})
}

//登录用户
func Login(c *gin.Context) {
	var data database.User
	var token string
	var message string
	//绑定数据
	c.ShouldBindJSON(&data)
	//验证登录
	loginCode := database.CheckLogin(data.Username, data.Password, data.Role)
	//此处设置token
	if loginCode == config.Success {
		token, _ = utils.GetToken(data.Username, data.Role)
		message = "登录成功"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"code":    loginCode,
			"token":   token,
		})
	} else {
		if loginCode == config.NoRole {
			message = "权限错误"
			c.JSON(http.StatusOK, gin.H{
				"message": message,
				"code":    loginCode,
			})
		}
		if loginCode == config.PasswordError {
			message = "密码错误"
			c.JSON(http.StatusOK, gin.H{
				"message": message,
				"code":    loginCode,
			})
		}
	}
}

//删除用户
func Delete(c *gin.Context) {
	var data database.User
	var message string
	c.ShouldBindJSON(&data)
	//验证当前用户权限
	if c.MustGet("role") != "admin" {
		message = "权限错误"
		c.JSON(http.StatusOK, gin.H{
			"code":    config.NoRole,
			"message": message,
		})
	} else {
		deleteUserCode, err := database.DeleteUser(data.Username)
		if deleteUserCode == config.Success {
			message = "删除成功"
			c.JSON(http.StatusOK, gin.H{
				"code":    deleteUserCode,
				"message": message,
			})
		}
		if deleteUserCode == config.DeleteUserError {
			message = "删除失败，请查看错误信息"
			c.JSON(http.StatusOK, gin.H{
				"code":    deleteUserCode,
				"message": message,
				"error":   err,
			})
		}
	}

}

//遍历所有用户信息
func AllUserInfo(c *gin.Context) {

}

//更新用户信息
func UpdateUserInfo(c *gin.Context) {

}

//根路由
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "訪問成功！",
	})
}
