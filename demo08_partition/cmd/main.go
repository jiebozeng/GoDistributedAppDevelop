package main

import (
	"GoDistributedAppDevelop/demo08_partition/config"
	"GoDistributedAppDevelop/demo08_partition/pkg/logs"
	"GoDistributedAppDevelop/demo08_partition/pkg/mysqldb"
	"GoDistributedAppDevelop/demo08_partition/router"
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
