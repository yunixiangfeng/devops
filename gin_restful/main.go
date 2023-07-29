package main

import "gin_restful/db"

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8806") // 启动服务了
}