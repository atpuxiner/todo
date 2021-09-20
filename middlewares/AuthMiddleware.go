package middlewares

import (
	"gitee.com/ax/todo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware()  gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// todo-完善
		if token :=ctx.GetHeader("Token");token == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": common.UNAUTHORIZED,
				"msg": common.CodeMsg(common.UNAUTHORIZED),
			})
			ctx.Abort()
		}else {
			ctx.Next()
		}
	}
}
