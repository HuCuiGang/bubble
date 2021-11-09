package main

import (

	"github.com/HuCuiGang/bubble/dao"
	"github.com/HuCuiGang/bubble/models"
	"github.com/HuCuiGang/bubble/routers"
)

func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出 关闭数据库连接
	models.InitModle() //模型绑定，自动建表
	//路由
	r := routers.SetupRouter()

	r.Run(":8080")
}