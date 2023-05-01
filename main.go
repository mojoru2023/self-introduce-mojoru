package main

import (
	"appbox_go_v/models"
	"appbox_go_v/routers"
	"appbox_go_v/utils"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE appbox;
	// 连接数据库
	err := utils.InitMySQL()
	if err != nil {
		panic(err)
	}

	defer utils.Close() // 程序退出关闭数据库连接

	// 模型绑定
	utils.DB.AutoMigrate(&models.OPENWORK_KAKAKU_JOBS{})
	// 注册路由

	r := routers.SetupRouter()
	r.Run(":8008")
}
