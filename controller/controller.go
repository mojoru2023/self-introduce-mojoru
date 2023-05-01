package controller

import (
	"appbox_go_v/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
    url     --> controller  --> logic   -->    model
   请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加

func CreateDt(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var dt models.OPENWORK_KAKAKU_JOBS
	c.BindJSON(&dt)
	// 2. 存入数据库
	err := models.CreateDt(&dt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		fmt.Println(dt)
		c.JSON(http.StatusOK, dt)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

// 查看所有的待办事项

func GetAllDt(c *gin.Context) {
	// 查询todo这个表里的所有数据
	dtList, err := models.GetAllDt()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, dtList)
	}

}

// 修改某一个待办事项

func UpdateOneDt(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	dt, err := models.GetOneDt(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&dt)
	if err = models.UpdateOneDt(dt); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, dt)
	}
}

// 删除某一个待办事项

func DeleteOneDt(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if _, err := models.DeleteOneDt(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})

	}

}
