package controllers

import (
	"gitee.com/ax/todo/common"
	"gitee.com/ax/todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查看列表
func GetTodoList(ctx *gin.Context) {
	var params models.ValidGetTodoList
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}
	err, data, total := models.GetTodoList(params)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data":nil,
		})
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  common.SUCCESS,
			"msg": common.CodeMsg(common.SUCCESS),
			"data":  data,
			"total": total,
		})
	}
}

// 创建数据
func CreateTodo(ctx *gin.Context) {
	var data models.Todo
	check := ctx.ShouldBindJSON(&data)
	if check != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  check.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}
	err := models.CreateTodo(data)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  common.SUCCESS,
			"msg": common.CodeMsg(common.SUCCESS),
			"data":  data,
		})
	}
}

// 删除数据
func DeleteTodo(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	err := models.DeleteTodo(tid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  common.SUCCESS,
			"msg": common.CodeMsg(common.SUCCESS),
			"data":  nil,
		})
	}
}

// 更新数据
func UpdateTodo(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	var data models.Todo
	check := ctx.ShouldBindJSON(&data)
	if check != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  check.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}
	err := models.UpdateTodo(tid, data)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  common.SUCCESS,
			"msg": common.CodeMsg(common.SUCCESS),
			"data":  data,
		})
	}
}

// 查询详情
func GetTodoDetail(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	err, data := models.GetTodoDetail(tid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.ERROR,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  common.SUCCESS,
			"msg": common.CodeMsg(common.SUCCESS),
			"data":  data,
		})
	}
}
