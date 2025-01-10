package main

import (
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/config"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/internal/produce_redpack"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/logs"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/pkg/redisdb"
	"GoDistributedAppDevelop/demo07_dstri_locks/redis_locks_server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置文件 viper
	config.Init()
	//初始化log
	logs.InitLogger(config.CONFIG.Logger.LogTypes, config.CONFIG.Logger.Dir, logs.LogEnvType(config.CONFIG.System.Mode), config.CONFIG.Logger.LogMaxAge)
	//初始化数据库链接
	mysqldb.InitMysql()
	//初始化Redis链接
	redisdb.Init()

	//启动协程 生成红包
	go func() {
		produce_redpack.ProduceRedpack()
	}()
	gin.SetMode(config.CONFIG.System.Mode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	router.InitRedpackRouter(r)
	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
}
