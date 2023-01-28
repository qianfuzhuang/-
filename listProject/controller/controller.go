package controller

import (
	"github.com/gin-gonic/gin"
	"listProject/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	//前端页面填写代办事项，点击提交，会发送请求到这里
	//1.从请求中吧数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateAtodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK,gin.H{
		//	"code":20000,
		//	"msg":"success",
		//	"data":todo,
		//})
	}
	//2.存入数据库
	//3.返回相应
}

func GetAlllist(c *gin.Context) {
	alllist, err := models.GetTodolist()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, alllist)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	todo, err := models.GetAtodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = models.Updata(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": ok,
		})
		return
	}
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}

}
