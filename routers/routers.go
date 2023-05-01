package routersvue2_highcharts

import (
	"appbox_go_v/controller"
	"appbox_go_v/middleware"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample example
// @host 127.0.0.1:8004
// @BasePath /api/v1

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	// router.Static("/static", "static")
	// // 告诉gin框架去哪里找模板文件
	// router.Use(gin.Recovery())
	router.Use(middleware.Options)
	// router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)

	//v1

	v1Group := router.Group("v1")
	{
		// 待办事项
		// 添加

		v1Group.POST("/vue2_highcharts/api", controller.CreateDt)
		v1Group.GET("/vue2_highcharts/api", controller.GetAllDt)
		v1Group.PUT("/vue2_highcharts/api/:id", controller.UpdateOneDt)
		v1Group.DELETE("/vue2_highcharts/api/:id", controller.DeleteOneDt)

	}

	return router

}
