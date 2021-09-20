package conf

import (
	"gitee.com/ax/todo/databases"
	"gitee.com/ax/todo/models"
	"github.com/spf13/viper"
	"os"
)

func DefaultInit() {
	initConf()
	databases.InitMysql()
	initMigrate()
}


func initConf() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")
	viper.AddConfigPath("/etc/backend")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func initMigrate() {
	_ = databases.DB.AutoMigrate(
		&models.Todo{},
	)
}
