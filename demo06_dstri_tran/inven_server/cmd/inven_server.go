package main

import (
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/config"
	"GoDistributedAppDevelop/demo06_dstri_tran/inven_server/internal/logic"
)
import "GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/logs"
import "GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/mysqldb"
import "GoDistributedAppDevelop/demo06_dstri_tran/inven_server/pkg/redisdb"

func main() {
	//初始化配置文件 viper
	config.Init()
	//初始化log
	logs.InitLogger(config.CONFIG.Logger.LogTypes, config.CONFIG.Logger.Dir, logs.LogEnvType(config.CONFIG.System.Mode), config.CONFIG.Logger.LogMaxAge)
	//初始化数据库链接
	mysqldb.InitMysql()
	//初始化Redis链接
	redisdb.Init()
	//库存订单逻辑
	invenOrderLgc := logic.OrderLgc{}
	invenOrderLgc.ListenPlaceOrder()
}
