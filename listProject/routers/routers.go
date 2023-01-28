package routers

import (
	"github.com/gin-gonic/gin"
	"listProject/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有
		v1Group.GET("/todo", controller.GetAlllist)
		//查看一个
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
