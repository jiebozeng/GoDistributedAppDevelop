package redisdb

import (
	"GoDistributedAppDevelop/demo06_dstri_tran/tran_server/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var RedisDb *redis.Client
var RedisCtx = context.Background()

func Init() {
	fmt.Println("redis init", config.CONFIG.Redis.Host, config.CONFIG.Redis.Port)
	// 从配置中读取对应的信息
	db, _ := strconv.Atoi(config.CONFIG.Redis.Database)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.CONFIG.Redis.Host + ":" + config.CONFIG.Redis.Port,
		Password: config.CONFIG.Redis.Password,
		DB:       db,
	})

	_, err := RedisDb.Ping(RedisCtx).Result()
	// todo 错误日志统一处理
	if err != nil {
		panic(err)
	}
}
