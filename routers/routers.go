package routers

import (
	"gitee.com/ax/todo/controllers"
	"gitee.com/ax/todo/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine{
	engine := gin.Default()
	engine.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is Todo.")
	})
	v1 := engine.Group("v1/todos")
	{
		v1.GET("", controllers.GetTodoList)
	}
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.POST("", controllers.CreateTodo)
		v1.PUT("/:tid", controllers.UpdateTodo)
		v1.DELETE("/:tid", controllers.DeleteTodo)
		v1.GET("/:tid", controllers.GetTodoDetail)
	}
	return engine
}
