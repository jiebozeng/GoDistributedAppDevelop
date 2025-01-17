package main

import (
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/config"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/internal/logics"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/logs"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/mysqldb"
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/pkg/redisdb"
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
	//创建一个订单
	orderLgc := logics.OrderLgc{}
	orderLgc.PlaceAnOrder(1, 1, 1)
}
