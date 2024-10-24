package controller

import (
	"Todolist/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateHandler(c *gin.Context) {
	var todo model.Todo
	//提取数据
	if err := c.ShouldBind(&todo); err != nil {
		panic(err)
		return
	}
	//存入数据库
	if err := model.CreateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		//返回响应
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	if todo, err := model.GetATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据库查找错误",
			"err": err.Error(),
		})
	} else {
		if err := model.DeleteATodo(todo); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "数据库删除错误",
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	if todo, err := model.GetATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据库查找错误",
			"err": err.Error(),
		})
	} else {
		c.ShouldBind(&todo)
		if err := model.UpdateATodo(todo); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "数据库更新错误",
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}

func ShowHandler(c *gin.Context) {
	if todos, err := model.GetTodolist(); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}
