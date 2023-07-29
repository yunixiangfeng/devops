package router

import (
	"github.com/gin-gonic/gin"
	. "gin_restful/api"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexUsers)

	// 路由群组
	users := router.Group("api/v1/users")
	{
		users.GET("", GetAll)
		users.POST("/add", AddUsers)
		users.GET("/get/:id", GetOne)
		users.POST("/update", UpdateUser)
		users.POST("/del", DelUser)
	}
	departments := router.Group("api/v1/department")
	{
		departments.GET("", GetAll)
		departments.POST("/add", AddUsers)
		departments.GET("/get/:id", GetOne)
		departments.POST("/update", UpdateUser)
		departments.POST("/del", DelUser)
	}
	return router
}
