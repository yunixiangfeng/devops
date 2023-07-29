package main

import (
	"gin_restful/db"

	"gin_restful/router"
)

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":8806") // 启动服务了
}
