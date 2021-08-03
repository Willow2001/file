package main

import (
	"owncloud/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"owncloud/controller"
)

func main() {
	r := gin.Default()
	err := dao.InitDB()//账号数据库
	if err !=nil{
		fmt.Print(err)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register",controller.UserRegister) //注册
		userGroup.POST("/login",controller.UserLogin)       //登录
		userGroup.GET("/logout",controller.UserLogout)      //注销
	}

	fileGroup := r.Group("/file")
	{
		fileGroup.POST("/upload",controller.FileUpload)       //上传单个文件
		fileGroup.POST("/multiupload",controller.MultiUpload) //上传多个文件
		fileGroup.POST("/download",controller.DownloadReadFile) //下载
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
