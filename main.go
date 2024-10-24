package main

import (
	"Todolist/dao"
	"Todolist/routers"
	_ "net/http"
)

func main() {

	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	r := routers.SetUpRouter()
	r.Run()

}
