package routers

import (
	"Todolist/controller"
	"Todolist/dao"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	v1Group := r.Group("/v1")
	{

		//增加
		v1Group.POST("/todo", controller.CreateHandler)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteHandler)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateHandler)
		//查看所有待办事项
		v1Group.GET("/todo", controller.ShowHandler)

	}

	return r
}
