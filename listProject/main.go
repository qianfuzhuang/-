package main

import (
	"listProject/dao"
	"listProject/models"
	"listProject/routers"
)

func main() {

	//连接数据库
	err := dao.InitmySQL()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":9091")

}
