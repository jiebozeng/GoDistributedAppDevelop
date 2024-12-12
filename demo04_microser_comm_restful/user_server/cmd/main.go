package main

import (
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置文件 viper
	config.Init()

	gin.SetMode(config.CONFIG.System.Mode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
}
