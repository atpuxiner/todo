package main

import (
	"gitee.com/ax/todo/conf"
	"gitee.com/ax/todo/routers"
	"github.com/spf13/viper"
)

func main() {
	conf.DefaultInit()
	r := routers.InitRouter()
	_ = r.Run(":" + viper.GetString("common.port"))
}
