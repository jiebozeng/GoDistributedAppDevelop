package main

import (
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/config"
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/pkg/logs"
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo04_microser_comm_restful/user_server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置文件 viper
	config.Init()
	//初始化log
	logs.InitLogger(config.CONFIG.Logger.LogTypes, config.CONFIG.Logger.Dir, logs.LogEnvType(config.CONFIG.System.Mode), config.CONFIG.Logger.LogMaxAge)
	//初始化数据库链接
	mysqldb.InitMysql()
	gin.SetMode(config.CONFIG.System.Mode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	router.InitUserRouter(r)
	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
}
